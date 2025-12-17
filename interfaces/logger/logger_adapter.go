// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier | MIT

// Package logger
package logger

import (
	"context"
	"fmt"
	"log/slog"
)

type Decorator struct {
	logger Interface
	prefix string
}

func NewLoggerAdapter(
	logger Interface,
	prefix string,
) *Decorator {
	return &Decorator{
		logger: logger,
		prefix: prefix,
	}
}

func (lg *Decorator) Init(logName string, logConfig *LogConfig) {
	lg.logger.Init(logName, logConfig)
}

func (lg *Decorator) Level() slog.Level {
	return lg.logger.Level()
}

func (lg *Decorator) ShutdownCallback(ctx context.Context) error {
	return lg.logger.ShutdownCallback(ctx)
}

func (lg *Decorator) LogHandler() *slog.Logger {
	return lg.logger.LogHandler()
}

func (lg *Decorator) Debug(msg string) {
	lg.logger.Debug(fmt.Sprintf("%s | %s", lg.prefix, msg))
}

func (lg *Decorator) Debugf(msg string, v ...interface{}) {
	lg.Debug(fmt.Sprintf(msg, v...))
}

func (lg *Decorator) Info(msg string) {
	lg.logger.Info(fmt.Sprintf("%s | %s", lg.prefix, msg))
}

func (lg *Decorator) Infof(msg string, v ...interface{}) {
	lg.Info(fmt.Sprintf(msg, v...))
}

func (lg *Decorator) Warn(msg string) {
	lg.logger.Warn(fmt.Sprintf("%s | %s", lg.prefix, msg))
}

func (lg *Decorator) Warnf(msg string, v ...interface{}) {
	lg.Warn(fmt.Sprintf(msg, v...))
}

func (lg *Decorator) Error(msg string) {
	lg.logger.Error(fmt.Sprintf("%s | %s", lg.prefix, msg))
}

func (lg *Decorator) Errorf(msg string, v ...interface{}) {
	lg.Error(fmt.Sprintf(msg, v...))
}

func (lg *Decorator) Fatal(msg string) {
	lg.logger.Fatal(fmt.Sprintf("%s | %s", lg.prefix, msg))
}

func (lg *Decorator) Fatalf(msg string, v ...interface{}) {
	lg.Debug(fmt.Sprintf(msg, v...))
}
