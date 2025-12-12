// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package testutils
package testutils

import "testing"

func TestNewFakeLogger(t *testing.T) {
	fakeLogger := NewFakeLogger(t)
	fakeLogger.Info("Test message")
}
