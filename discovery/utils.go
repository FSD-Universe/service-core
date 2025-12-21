// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package discovery
package discovery

import (
	"context"
	"slices"
	"time"

	"half-nothing.cn/service-core/interfaces/cleaner"
	"half-nothing.cn/service-core/interfaces/discovery"
	"half-nothing.cn/service-core/interfaces/logger"
	"half-nothing.cn/service-core/utils"
)

func StartServiceDiscovery(
	ctx context.Context,
	lg logger.Interface,
	cl cleaner.Interface,
	started chan bool,
	version *utils.Version,
	serviceName string,
	servicePort int,
) *ServiceDiscovery {
	start := <-started
	if !start {
		return nil
	}
	service := NewServiceDiscovery(lg, serviceName, servicePort, version)
	if err := service.Start(ctx); err != nil {
		lg.Fatalf("fail to start service discovery: %v", err)
		return nil
	}
	cl.Add("Service Discovery", service.Stop)
	return service
}

func StartServiceListener(
	ctx context.Context,
	cl cleaner.Interface,
	handler func(status *ServiceStatusChange),
) *ServiceListener {
	listener := NewServiceListener(nil, handler)
	listener.Start(ctx)
	cl.Add("Service Listener", listener.Stop)
	return listener
}

func KeepRequiredServiceOnline(
	logger logger.Interface,
	requiredServices []string,
	service *ServiceDiscovery,
	timeout time.Duration,
	shutdown func(),
	flushService func(info *discovery.ServiceInfo),
) func(status *ServiceStatusChange) {
	return func(status *ServiceStatusChange) {
		if status.Status != ServiceUnregistered {
			return
		}
		if slices.Contains(requiredServices, status.ServiceName) {
			logger.Warnf("required service %s offline, wait %.0fs for online", status.ServiceName, timeout.Seconds())
			info, err := service.WaitForService(status.ServiceName, timeout)
			if err != nil {
				logger.Errorf("required service %s offline, shutting down", status.ServiceName)
				shutdown()
				return
			}
			logger.Infof("required service %s online: %s:%d", info.Name, info.IP, info.Port)
			flushService(info)
		}
	}
}
