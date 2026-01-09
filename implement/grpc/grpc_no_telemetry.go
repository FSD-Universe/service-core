// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package grpc
package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"half-nothing.cn/service-core/interfaces/cleaner"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/logger"
)

func StartGrpcServer(
	lg logger.Interface,
	cl cleaner.Interface,
	c *config.GrpcServerConfig,
	started chan bool,
	initServer func(s *grpc.Server),
) {
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		lg.Fatalf("gRPC fail to listen: %v", err)
		started <- false
		return
	}
	var s *grpc.Server
	if c.TLSConfig.Enable {
		s = grpc.NewServer(
			grpc.Creds(c.TLSConfig.Credentials),
		)
	} else {
		s = grpc.NewServer()
	}
	initServer(s)
	reflection.Register(s)
	cl.Add("gRPC Server", func(ctx context.Context) error {
		close(started)
		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		cleanOver := make(chan struct{})
		go func() {
			s.GracefulStop()
			cleanOver <- struct{}{}
		}()
		select {
		case <-timeoutCtx.Done():
			s.Stop()
		case <-cleanOver:
		}
		return nil
	})
	lg.Infof("gRPC server listening at %v", lis.Addr())
	started <- true
	if err := s.Serve(lis); err != nil {
		lg.Fatalf("gRPC failed to serve: %v", err)
		return
	}
}

func StartGrpcClient(
	lg logger.Interface,
	host string,
	port int,
	c *config.GrpcClientConfig,
) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var creds credentials.TransportCredentials
	if c.EnableTLS {
		creds = c.Credentials
	} else {
		creds = insecure.NewCredentials()
	}
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		lg.Fatalf("gRPC client connect fail: %v", err)
		return nil, err
	}
	return conn, nil
}
