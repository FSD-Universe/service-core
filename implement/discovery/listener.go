// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package discovery
package discovery

import (
	"context"

	capi "github.com/hashicorp/consul/api"
	"half-nothing.cn/service-core/interfaces/discovery"
	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/interfaces/logger"
)

type ServiceListener struct {
	serviceStatusChan chan *discovery.ServiceEvent
	handler           func(status *discovery.ServiceEvent)
	ctx               context.Context
	cancel            context.CancelFunc
}

func NewServiceListener(
	statusChan chan *discovery.ServiceEvent,
	handler func(status *discovery.ServiceEvent),
) *ServiceListener {
	return &ServiceListener{
		serviceStatusChan: statusChan,
		handler:           handler,
	}
}

func (sl *ServiceListener) Start(ctx context.Context) {
	sl.ctx, sl.cancel = context.WithCancel(ctx)
	go sl.listener()
}

func (sl *ServiceListener) Stop(_ context.Context) error {
	sl.cancel()
	return nil
}

func (sl *ServiceListener) listener() {
	for {
		select {
		case status := <-sl.serviceStatusChan:
			go sl.handler(status)
		case <-sl.ctx.Done():
			return
		}
	}
}

func KeepRequiredServiceOnline(
	logger logger.Interface,
	service discovery.ServiceManagerInterface,
	shutdown func(),
	flushService func(serviceName string, info *capi.ServiceEntry),
) func(status *discovery.ServiceEvent) {
	return func(status *discovery.ServiceEvent) {
		logger.Debugf("received service event: %s %s", status.ServiceName, status.EventType)
		switch status.EventType {
		case discovery.ServiceUpdate:
			fallthrough
		case discovery.ServiceOnline:
			serviceInfo := service.GetRandomServiceInfo(status.ServiceName)
			flushService(status.ServiceName, serviceInfo)
		case discovery.ServiceOffline:
			timeout := *global.ReconnectTimeout
			logger.Warnf("required service %s offline, wait %.0fs for online", status.ServiceName, timeout.Seconds())
			if err := service.WaitForServices(timeout); err != nil {
				logger.Errorf("required service %s offline, shutting down", status.ServiceName)
				shutdown()
				return
			}
			serviceInfo := service.GetRandomServiceInfo(status.ServiceName)
			logger.Infof("required service %s online: %s:%d", status.ServiceName, serviceInfo.Service.Address, serviceInfo.Service.Port)
			flushService(status.ServiceName, serviceInfo)
		default:
			return
		}
	}
}
