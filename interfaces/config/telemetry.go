// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build telemetry

// Package config
package config

import "fmt"

type TelemetryConfig struct {
	Enable          bool   `yaml:"enable"`
	Name            string `yaml:"name"`
	Endpoint        string `yaml:"endpoint"`
	Token           string `yaml:"token"`
	ServiceName     string `yaml:"service_name"`
	HostName        string `yaml:"host_name"`
	DatabaseTrace   bool   `yaml:"database_trace"`
	GrpcServerTrace bool   `yaml:"grpc_server_trace"`
	GrpcClientTrace bool   `yaml:"grpc_client_trace"`
	HttpServerTrace bool   `yaml:"http_server_trace"`
}

func (t *TelemetryConfig) InitDefaults() {
	t.Enable = false
	t.Name = "service"
	t.Endpoint = "http://localhost:4318"
	t.Token = ""
	t.ServiceName = "service"
	t.HostName = "localhost"
	t.DatabaseTrace = false
	t.GrpcServerTrace = false
	t.GrpcClientTrace = false
	t.HttpServerTrace = false
}

func (t *TelemetryConfig) Verify() (bool, error) {
	if !t.Enable {
		t.InitDefaults()
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
