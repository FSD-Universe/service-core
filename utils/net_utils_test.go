// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "testing"

func TestGetLocalIP(t *testing.T) {
	t.Log(GetLocalIP("WLAN"))
}

func TestGetBroadcastIP(t *testing.T) {
	localIP := GetLocalIP("WLAN")
	t.Log(localIP)
	t.Log(GetBroadcastIP("WLAN", localIP))
}
