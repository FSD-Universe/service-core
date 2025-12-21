// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package grpc
package grpc

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"half-nothing.cn/service-core/interfaces/logger"
)

type ClientConnections struct {
	logger      logger.Interface
	connections map[string]*grpc.ClientConn
	lock        sync.RWMutex
}

func NewClientConnections(lg logger.Interface) *ClientConnections {
	return &ClientConnections{
		logger:      logger.NewLoggerAdapter(lg, "grpc-connections"),
		connections: make(map[string]*grpc.ClientConn),
		lock:        sync.RWMutex{},
	}
}

func (c *ClientConnections) Get(name string) *grpc.ClientConn {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if conn, ok := c.connections[name]; ok {
		return conn
	}
	return nil
}

func (c *ClientConnections) Add(name string, conn *grpc.ClientConn) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.connections[name] = conn
}

func (c *ClientConnections) Remove(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.connections, name)
}

func (c *ClientConnections) Close(ctx context.Context) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	eg, _ := errgroup.WithContext(ctx)
	for name, conn := range c.connections {
		eg.Go(func() error {
			c.logger.Debugf("closing connection: %s", name)
			return conn.Close()
		})
	}
	return eg.Wait()
}
