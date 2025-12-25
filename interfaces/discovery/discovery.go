// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package discovery
package discovery

import (
	"context"
	"time"

	capi "github.com/hashicorp/consul/api"
)

type EventType string

const (
	ServiceError   EventType = "error"
	ServiceUpdate  EventType = "update"
	ServiceOnline  EventType = "online"
	ServiceOffline EventType = "offline"
)

type ServiceEvent struct {
	ServiceName string
	Instances   []*capi.ServiceEntry
	EventType   EventType
}

type Interface interface {
	RegisterServer() error
	UnregisterServer(ctx context.Context) error
	QueryService(serviceName string, queryOptions *capi.QueryOptions) ([]*capi.ServiceEntry, *capi.QueryMeta, error)
}

type ServiceManagerInterface interface {
	WatchService(serviceName string)
	WatchServices(serviceNames []string)
	GetServiceState(serviceName string) []*capi.ServiceEntry
	GetRandomServiceInfo(serviceName string) *capi.ServiceEntry
	CheckHealthy() bool
	StopWatch(ctx context.Context) error
	WaitForServices(timeout time.Duration) error
}
