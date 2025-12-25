// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build !telemetry

// Package grpc
package grpc

import (
	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/logger"
)

func InitGrpcClient(lg *logger.Logger, clientConfig *config.GrpcClientConfig, info *capi.ServiceEntry) (conn *grpc.ClientConn, err error) {
	conn, err = StartGrpcClient(lg, info.Service.Address, info.Service.Port, clientConfig)
	if err != nil {
		lg.Fatalf("fail to get grpc client connection: %v", err)
	}
	return
}
