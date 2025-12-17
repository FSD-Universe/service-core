// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build grpc && telemetry

// Package grpc
package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"half-nothing.cn/service-core/interfaces/cleaner"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/logger"
)

func StartGrpcServerWithTrace(
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
		return
	}
	var s *grpc.Server
	if c.TLSConfig.Enable {
		s = grpc.NewServer(
			grpc.Creds(c.TLSConfig.Credentials),
			grpc.StatsHandler(otelgrpc.NewServerHandler()),
		)
	} else {
		s = grpc.NewServer(
			grpc.StatsHandler(otelgrpc.NewServerHandler()),
		)
	}
	initServer(s)
	reflection.Register(s)
	cl.Add("gRPC Server", func(ctx context.Context) error {
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
	close(started)
	if err := s.Serve(lis); err != nil {
		lg.Fatalf("gRPC failed to serve: %v", err)
		return
	}
}

func StartGrpcClientWithTrace(
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
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		lg.Fatalf("gRPC client connect fail: %v", err)
		return nil, err
	}
	return conn, nil
}
