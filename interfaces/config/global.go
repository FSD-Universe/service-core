// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import (
	"fmt"

	"half-nothing.cn/service-core/interfaces/logger"
	"half-nothing.cn/service-core/utils"
)

type GlobalConfig struct {
	Name      string            `yaml:"name"`
	Version   string            `yaml:"version"`
	LogConfig *logger.LogConfig `yaml:"log"`

	// 内部使用
	ConfigVersion string `yaml:"-"`
}

func (g *GlobalConfig) InitDefaults() {
	g.LogConfig = &logger.LogConfig{}
	g.LogConfig.InitDefaults()
}

func (g *GlobalConfig) Verify() (bool, error) {
	if g.Name == "" {
		return false, fmt.Errorf("global name is empty")
	}
	if g.Version == "" {
		return false, fmt.Errorf("global version is empty")
	}
	configVersion := utils.NewVersion(g.Version)
	if configVersion == nil {
		return false, fmt.Errorf("global version is invalid: %s", g.Version)
	}
	targetConfigVersion := utils.NewVersion(g.ConfigVersion)
	if targetConfigVersion.CheckVersion(configVersion) != utils.AllMatch {
		return false, fmt.Errorf("config version mismatch, expected %s, got %s", g.ConfigVersion, g.Version)
	}
	if g.LogConfig == nil {
		return false, fmt.Errorf("log config is empty")
	}
	if ok, err := g.LogConfig.Verify(); !ok {
		return false, err
	}
	return true, nil
}
