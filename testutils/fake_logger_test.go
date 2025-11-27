// Package testutils
package testutils

import "testing"

func TestNewFakeLogger(t *testing.T) {
	fakeLogger := NewFakeLogger(t)
	fakeLogger.Info("Test message")
}
