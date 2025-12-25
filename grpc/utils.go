// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build telemetry

// Package grpc
package grpc

import (
	"google.golang.org/grpc"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/logger"
)

func InitGrpcClient(lg *logger.Logger, teleConfig *config.TelemetryConfig, clientConfig *config.GrpcClientConfig, host string, port int) (conn *grpc.ClientConn, err error) {
	if teleConfig.GrpcClientTrace {
		conn, err = StartGrpcClientWithTrace(lg, host, port, clientConfig)
	} else {
		conn, err = StartGrpcClient(lg, host, port, clientConfig)
	}
	if err != nil {
		lg.Fatalf("fail to get grpc client connection: %v", err)
	}
	return
}
