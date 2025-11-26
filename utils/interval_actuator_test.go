// Package utils
package utils

import (
	"testing"
	"time"
)

func TestNewIntervalActuator(t *testing.T) {
	actuator := NewIntervalActuator(time.Second, func() {
		t.Log("callback")
	})
	actuator.Start()
	time.Sleep(time.Second * 2)
	actuator.Stop()
}
