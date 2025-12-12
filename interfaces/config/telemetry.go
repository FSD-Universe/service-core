// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build telemetry

// Package config
package config

import "fmt"

type TelemetryConfig struct {
	Enable           bool   `yaml:"enable"`
	Name             string `yaml:"name"`
	Endpoint         string `yaml:"endpoint"`
	Token            string `yaml:"token"`
	ServiceName      string `yaml:"service_name"`
	HostName         string `yaml:"host_name"`
	DatabaseTracking bool   `yaml:"database_tracking"`
}

func (t *TelemetryConfig) InitDefaults() {
	t.Enable = false
	t.Name = "service"
	t.Endpoint = "http://localhost:4318"
	t.Token = ""
	t.ServiceName = "service"
	t.HostName = "localhost"
	t.DatabaseTracking = false
}

func (t *TelemetryConfig) Verify() (bool, error) {
	if !t.Enable {
		return true, nil
	}
	if t.Name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}
	if t.Endpoint == "" {
		return false, fmt.Errorf("endpoint cannot be empty")
	}
	if t.Token == "" {
		return false, fmt.Errorf("token cannot be empty")
	}
	if t.ServiceName == "" {
		return false, fmt.Errorf("service name cannot be empty")
	}
	if t.HostName == "" {
		return false, fmt.Errorf("host name cannot be empty")
	}
	return true, nil
}
