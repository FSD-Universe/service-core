// Package global
package global

import (
	"os"
	"strconv"
	"time"

	"half-nothing.cn/service-core/utils"
)

func CheckBoolEnv(envKey string, target *bool) {
	value := os.Getenv(envKey)
	if val, err := strconv.ParseBool(value); err == nil && val {
		*target = true
	}
}

func CheckStringEnv(envKey string, target *string) {
	value := os.Getenv(envKey)
	if value != "" {
		*target = value
	}
}

func CheckIntEnv(envKey string, target *int, defaultValue int) {
	value := os.Getenv(envKey)
	if value != "" {
		*target = utils.StrToInt(value, defaultValue)
	}
}

func CheckDurationEnv(envKey string, target *time.Duration) {
	value := os.Getenv(envKey)
	if duration, err := time.ParseDuration(value); err == nil {
		*target = duration
	}
}
