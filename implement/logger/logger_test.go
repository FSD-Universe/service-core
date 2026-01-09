// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package logger
package logger

import (
	"testing"

	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/interfaces/logger"
)

func TestNewLogger(t *testing.T) {
	lg := NewLogger()
	*global.NoLogs = true
	logConfig := &logger.LogConfig{}
	logConfig.InitDefaults()
	lg.Init("MAIN", logConfig)
	lg.Debug("This is a debug message")
	lg.Info("This is an info message")
	lg.Warn("This is a warning message")
	lg.Error("This is an error message")
}
