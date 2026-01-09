// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package logger
package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"half-nothing.cn/service-core/interfaces/global"
)

type LogConfig struct {
	Level      string `yaml:"level"`
	Path       string `yaml:"path"`
	Rotate     bool   `yaml:"rotate"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
	LocalTime  bool   `yaml:"local_time"`
}

var levels = []string{"debug", "info", "warn", "error", "fatal"}

func (l *LogConfig) InitDefaults() {
	l.Level = "info"
	l.Path = "./logs/logs.log"
	l.Rotate = true
	l.MaxSize = 2
	l.MaxAge = 28
	l.MaxBackups = 30
	l.Compress = true
	l.LocalTime = true
}

func (l *LogConfig) Verify() (bool, error) {
	if l.Level == "" {
		return false, fmt.Errorf("log level is empty")
	}
	level := strings.ToLower(l.Level)
	if !slices.Contains(levels, level) {
		return false, fmt.Errorf("log level is invalid")
	}
	if l.Path == "" {
		return false, fmt.Errorf("log path is empty")
	}
	if err := os.MkdirAll(filepath.Dir(l.Path), global.DefaultDirectoryPermission); err != nil {
		return false, fmt.Errorf("create log directory failed: %s", err)
	}
	if l.Rotate {
		if l.MaxSize <= 0 {
			return false, fmt.Errorf("log max size is invalid")
		}
		if l.MaxAge <= 0 {
			return false, fmt.Errorf("log max age is invalid")
		}
		if l.MaxBackups <= 0 {
			return false, fmt.Errorf("log max backups is invalid")
		}
	}
	return true, nil
}
