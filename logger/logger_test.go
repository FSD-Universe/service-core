// Package logger
package logger

import (
	"testing"

	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/global"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	*global.NoLogs = true
	logConfig := &config.LogConfig{}
	logConfig.InitDefaults()
	logger.Init("MAIN", logConfig)
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
