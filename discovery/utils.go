// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package discovery
package discovery

import (
	"half-nothing.cn/service-core/interfaces/cleaner"
	"half-nothing.cn/service-core/interfaces/logger"
	"half-nothing.cn/service-core/utils"
)

func StartServiceDiscovery(lg logger.Interface, cl cleaner.Interface, started chan bool, version *utils.Version, serviceName string, servicePort int) {
	start := <-started
	if !start {
		return
	}
	service := NewServiceDiscovery(lg, serviceName, servicePort, version)
	if err := service.Start(); err != nil {
		lg.Fatalf("fail to start service discovery: %v", err)
		return
	}
	cl.Add("Service Discovery", service.Stop)
}
