// Package logger
package logger

import (
	"testing"

	"github.com/FSD-Universe/service-core/interfaces/config"
	"github.com/FSD-Universe/service-core/interfaces/global"
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
