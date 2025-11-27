// Package testutils
package testutils

import (
	"context"
	"log/slog"
	"testing"

	"github.com/FSD-Universe/service-core/interfaces/config"
)

type FakeLogger struct {
	t *testing.T
}

func NewFakeLogger(t *testing.T) *FakeLogger {
	return &FakeLogger{
		t: t,
	}
}

func (f *FakeLogger) Init(_ string, _ *config.LogConfig) {
	return
}

func (f *FakeLogger) ShutdownCallback(_ context.Context) error {
	return nil
}

func (f *FakeLogger) LogHandler() *slog.Logger {
	return nil
}

func (f *FakeLogger) Debug(msg string) {
	f.t.Logf(msg)
}

func (f *FakeLogger) Debugf(msg string, v ...interface{}) {
	f.t.Logf(msg, v...)
}

func (f *FakeLogger) Info(msg string) {
	f.t.Log(msg)
}

func (f *FakeLogger) Infof(msg string, v ...interface{}) {
	f.t.Logf(msg, v...)
}

func (f *FakeLogger) Warn(msg string) {
	f.t.Log(msg)
}

func (f *FakeLogger) Warnf(msg string, v ...interface{}) {
	f.t.Logf(msg, v...)
}

func (f *FakeLogger) Error(msg string) {
	f.t.Log(msg)
}

func (f *FakeLogger) Errorf(msg string, v ...interface{}) {
	f.t.Logf(msg, v...)
}

func (f *FakeLogger) Fatal(msg string) {
	f.t.Fatal(msg)
}

func (f *FakeLogger) Fatalf(msg string, v ...interface{}) {
	f.t.Fatalf(msg, v...)
}
