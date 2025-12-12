// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/global"
)

// readConfig 读取并解析配置文件
func readConfig[T config.DefaultSetter]() (T, error) {
	var c T
	if reflect.ValueOf(c).Kind() == reflect.Pointer {
		t := reflect.TypeOf(c)
		if t.Elem().Kind() == reflect.Struct {
			newValue := reflect.New(t.Elem())
			c = newValue.Interface().(T)
		}
	}
	c.InitDefaults()

	file, err := os.OpenFile(*global.ConfigFilePath, os.O_RDONLY, global.DefaultFilePermissions)
	defer func(file *os.File) { _ = file.Close() }(file)
	var zero T

	if err != nil {
		if err := saveConfig(c); err != nil {
			return zero, fmt.Errorf("fail to save configuration file while creating configuration file, %s", err)
		} else {
			return c, nil
		}
	}

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(c); err != nil {
		return zero, err
	}

	return c, nil
}

// saveConfig 将配置保存到配置文件
func saveConfig[T config.DefaultSetter](c T) error {
	file, err := os.OpenFile(*global.ConfigFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, global.DefaultFilePermissions)
	defer func(file *os.File) { _ = file.Close() }(file)
	if err != nil {
		return err
	}

	encoder := yaml.NewEncoder(file)
	defer func(encoder *yaml.Encoder) { _ = encoder.Close() }(encoder)
	return encoder.Encode(c)
}

// Manager 配置管理器结构体
type Manager[T config.Item] struct {
	config T
}

// NewManager 创建一个新的配置管理器实例
func NewManager[T config.Item]() *Manager[T] {
	return &Manager[T]{}
}

func (m *Manager[T]) Init() error {
	c, err := readConfig[T]()
	if err != nil {
		return err
	}
	if ok, err := c.Verify(); !ok {
		return err
	}
	m.config = c
	return nil
}

func (m *Manager[T]) GetConfig() T {
	return m.config
}

func (m *Manager[T]) SaveConfig() error {
	return saveConfig(m.config)
}
