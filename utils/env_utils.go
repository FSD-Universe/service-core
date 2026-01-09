// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import (
	"os"
	"strconv"
	"time"
)

func CheckBoolEnv(envKey string, target *bool) {
	value := os.Getenv(envKey)
	if val, err := strconv.ParseBool(value); err == nil {
		*target = val
	}
}

func CheckStringEnv(envKey string, target *string) {
	val := os.Getenv(envKey)
	if val != "" && *target != val {
		*target = val
	}
}

func CheckIntEnv(envKey string, target *int) {
	value := StrToInt(os.Getenv(envKey), *target)
	if *target != value {
		*target = value
	}
}

func CheckDurationEnv(envKey string, target *time.Duration) {
	value := os.Getenv(envKey)
	if duration, err := time.ParseDuration(value); err == nil {
		*target = duration
	}
}
