// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "testing"

func TestGetLocalIP(t *testing.T) {
	t.Log(GetLocalIP("WLAN"))
}
