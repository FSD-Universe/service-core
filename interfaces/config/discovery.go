// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import "fmt"

type DiscoveryConfig struct {
	Id       string `yaml:"id"`
	Name     string `yaml:"name"`
	HttpPort int    `yaml:"http_port"`
	GrpcPort int    `yaml:"grpc_port"`
}

func (d *DiscoveryConfig) InitDefaults() {
	d.Id = "service-1"
	d.Name = "service"
	d.HttpPort = 8080
	d.GrpcPort = 8081
}

func (d *DiscoveryConfig) Verify() (bool, error) {
	if d.Id == "" {
		return false, fmt.Errorf("service id is empty")
	}
	if d.Name == "" {
		return false, fmt.Errorf("service name is empty")
	}
	if d.HttpPort <= 0 || d.HttpPort > 65535 {
		return false, fmt.Errorf("http port is invalid")
	}
	if d.GrpcPort <= 0 || d.GrpcPort > 65535 {
		return false, fmt.Errorf("grpc port is invalid")
	}
	if d.HttpPort == d.GrpcPort {
		return false, fmt.Errorf("http port and grpc port cannot be the same")
	}
	return true, nil
}
