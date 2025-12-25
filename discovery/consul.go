// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package discovery
package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	capi "github.com/hashicorp/consul/api"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/discovery"
	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/interfaces/logger"
)

type ConsulClient struct {
	logger           logger.Interface
	config           *config.DiscoveryConfig
	version          string
	client           *capi.Client
	serviceStates    map[string][]*capi.ServiceEntry
	stateMutex       sync.RWMutex
	EventChan        chan *discovery.ServiceEvent
	watchCancelFuncs map[string]context.CancelFunc
}

func NewConsulClient(
	lg logger.Interface,
	c *config.DiscoveryConfig,
	version string,
) *ConsulClient {
	return &ConsulClient{
		logger:  logger.NewLoggerAdapter(lg, "consul-client"),
		config:  c,
		version: version,
		client:  nil,
	}
}

func (consul *ConsulClient) RegisterServer() error {
	c := capi.DefaultConfig()
	c.Address = *global.CenterAddress
	var err error
	consul.client, err = capi.NewClient(c)
	if err != nil {
		consul.logger.Fatalf("consul client init fail: %v", err)
		return err
	}
	consul.logger.Infof("consul client init success")
	registration := &capi.AgentServiceRegistration{
		ID:      consul.config.Id,
		Name:    consul.config.Name,
		Address: *global.ServiceAddress,
		Port:    consul.config.GrpcPort,
		Check: &capi.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", *global.ServiceAddress, consul.config.HttpPort),
			Interval:                       *global.HealthCheckInterval,
			Timeout:                        *global.HealthCheckTimeout,
			DeregisterCriticalServiceAfter: *global.DeregisterAfter,
		},
		Meta: map[string]string{"version": consul.version},
	}
	err = consul.client.Agent().ServiceRegister(registration)
	if err != nil {
		consul.logger.Fatalf("register server fail: %v", err)
		return err
	}
	consul.logger.Infof("register server success")
	return nil
}

func (consul *ConsulClient) UnregisterServer(ctx context.Context) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	errChan := make(chan error)
	defer close(errChan)
	go func() {
		errChan <- consul.client.Agent().ServiceDeregister(consul.config.Id)
	}()
	select {
	case <-timeoutCtx.Done():
		consul.logger.Errorf("timeout while deregistering server")
		return timeoutCtx.Err()
	case err := <-errChan:
		return err
	}
}

func (consul *ConsulClient) QueryService(serviceName string, queryOptions *capi.QueryOptions) ([]*capi.ServiceEntry, *capi.QueryMeta, error) {
	serviceEntries, meta, err := consul.client.Health().Service(serviceName, "", true, queryOptions)
	if err != nil {
		consul.logger.Errorf("query service fail: %v", err)
		return nil, nil, err
	}
	if len(serviceEntries) == 0 {
		consul.logger.Errorf("no service for %s", serviceName)
		return nil, nil, fmt.Errorf("no server for %s", serviceName)
	}
	consul.logger.Debugf("total %d services for %s", len(serviceEntries), serviceName)
	return serviceEntries, meta, nil
}

func (consul *ConsulClient) WatchService(serviceName string) {
	ctx, cancel := context.WithCancel(context.Background())
	consul.watchCancelFuncs[serviceName] = cancel

	lastIndex := uint64(0)
	lastEntries := make([]*capi.ServiceEntry, 0)
	go func(name string) {
		for {
			select {
			case <-ctx.Done():
				consul.logger.Debugf("cancel watch service %s", serviceName)
				return
			default:
				entries, meta, err := consul.client.Health().Service(
					name, "", true,
					&capi.QueryOptions{
						WaitIndex: lastIndex,
						WaitTime:  30 * time.Second,
					},
				)
				if err != nil {
					consul.EventChan <- &discovery.ServiceEvent{
						ServiceName: name,
						EventType:   discovery.ServiceError,
					}
					time.Sleep(5 * time.Second)
					continue
				}
				if meta.LastIndex <= lastIndex {
					continue
				}
				lastIndex = meta.LastIndex
				consul.stateMutex.Lock()
				consul.serviceStates[name] = entries
				consul.stateMutex.Unlock()
				if len(entries) == 0 {
					consul.EventChan <- &discovery.ServiceEvent{
						ServiceName: name,
						EventType:   discovery.ServiceOffline,
					}
				} else if len(lastEntries) == 0 {
					consul.EventChan <- &discovery.ServiceEvent{
						ServiceName: name,
						Instances:   entries,
						EventType:   discovery.ServiceOnline,
					}
				} else {
					consul.EventChan <- &discovery.ServiceEvent{
						ServiceName: name,
						Instances:   entries,
						EventType:   discovery.ServiceUpdate,
					}
				}
				lastEntries = entries
			}
		}
	}(serviceName)
}

func (consul *ConsulClient) WatchServices(serviceNames []string) {
	for _, name := range serviceNames {
		consul.WatchService(name)
	}
}

func (consul *ConsulClient) GetServiceState(serviceName string) []*capi.ServiceEntry {
	consul.stateMutex.RLock()
	defer consul.stateMutex.RUnlock()
	return consul.serviceStates[serviceName]
}

func (consul *ConsulClient) GetRandomServiceInfo(serviceName string) *capi.ServiceEntry {
	consul.stateMutex.RLock()
	defer consul.stateMutex.RUnlock()
	services := consul.serviceStates[serviceName]
	return services[rand.Intn(len(services))]
}

func (consul *ConsulClient) CheckHealthy() bool {
	consul.stateMutex.RLock()
	defer consul.stateMutex.RUnlock()
	for _, instances := range consul.serviceStates {
		if len(instances) == 0 {
			return false
		}
	}
	return true
}

func (consul *ConsulClient) StopWatch(_ context.Context) error {
	for _, cancel := range consul.watchCancelFuncs {
		cancel()
	}
	return nil
}

func (consul *ConsulClient) WaitForServices(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if consul.CheckHealthy() {
				return nil
			}
			time.Sleep(1 * time.Second)
		}
	}
}
