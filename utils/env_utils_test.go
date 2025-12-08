// Package utils
package utils

import (
	"os"
	"testing"
	"time"
)

func TestCheckBoolEnv(t *testing.T) {
	envKey := "TEST_BOOL"
	target := false
	_ = os.Setenv(envKey, "true")
	CheckBoolEnv(envKey, &target)
	if target != true {
		t.Errorf("CheckBoolEnv() = %v, want %v", target, true)
	}
	_ = os.Setenv(envKey, "false")
	CheckBoolEnv(envKey, &target)
	if target != false {
		t.Errorf("CheckBoolEnv() = %v, want %v", target, false)
	}
	_ = os.Setenv(envKey, "1")
	CheckBoolEnv(envKey, &target)
	if target != true {
		t.Errorf("CheckBoolEnv() = %v, want %v", target, true)
	}
	_ = os.Setenv(envKey, "0")
	CheckBoolEnv(envKey, &target)
	if target != false {
		t.Errorf("CheckBoolEnv() = %v, want %v", target, false)
	}
	_ = os.Unsetenv(envKey)
}

func TestCheckStringEnv(t *testing.T) {
	envKey := "TEST_STRING"
	target := ""
	_ = os.Setenv(envKey, "test")
	CheckStringEnv(envKey, &target)
	if target != "test" {
		t.Errorf("CheckStringEnv() = %v, want %v", target, "test")
	}
	target = "test2"
	_ = os.Unsetenv(envKey)
	CheckStringEnv(envKey, &target)
	if target != "test2" {
		t.Errorf("CheckStringEnv() = %v, want %v", target, "test2")
	}
}

func TestCheckDurationEnv(t *testing.T) {
	envKey := "TEST_DURATION"
	target := time.Duration(0)
	_ = os.Setenv(envKey, "1s")
	CheckDurationEnv(envKey, &target)
	if target != time.Second {
		t.Errorf("CheckDurationEnv() = %v, want %v", target, time.Second)
	}
	_ = os.Setenv(envKey, "1m")
	CheckDurationEnv(envKey, &target)
	if target != time.Minute {
		t.Errorf("CheckDurationEnv() = %v, want %v", target, time.Minute)
	}
	_ = os.Setenv(envKey, "1d")
	CheckDurationEnv(envKey, &target)
	if target != time.Minute {
		t.Errorf("CheckDurationEnv() = %v, want %v", target, time.Minute)
	}
	_ = os.Unsetenv(envKey)
}

func TestCheckIntEnv(t *testing.T) {
	envKey := "TEST_INT"
	target := 0
	_ = os.Setenv(envKey, "1")
	CheckIntEnv(envKey, &target)
	if target != 1 {
		t.Errorf("CheckIntEnv() = %v, want %v", target, 1)
	}
	_ = os.Setenv(envKey, "0")
	CheckIntEnv(envKey, &target)
	if target != 0 {
		t.Errorf("CheckIntEnv() = %v, want %v", target, 0)
	}
	_ = os.Setenv(envKey, "aaa")
	target = 1
	CheckIntEnv(envKey, &target)
	if target != 1 {
		t.Errorf("CheckIntEnv() = %v, want %v", target, 1)
	}
	_ = os.Unsetenv(envKey)
}
