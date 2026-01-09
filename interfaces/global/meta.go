// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package global
package global

import (
	"flag"
	"time"

	"half-nothing.cn/service-core/utils"
)

var (
	NoLogs         = flag.Bool("no_logs", false, "Disable logging to file")
	AutoMigrate    = flag.Bool("auto_migrate", false, "Auto migrate database. Caution: Dont set this to true in production env")
	ConfigFilePath = flag.String("config", "./config.yaml", "Path to configuration file")
)

// 服务发现配置
var (
	HealthCheckInterval = flag.String("health_check_interval", "30s", "Health check interval")
	HealthCheckTimeout  = flag.String("health_check_timeout", "5s", "Health check timeout")
	DeregisterAfter     = flag.String("deregister_after", "1m", "Deregister critical service after")
	ServiceAddress      = flag.String("service_address", "localhost", "Service address")
	CenterAddress       = flag.String("center_address", "localhost:8500", "Service discovery center address")
	ReconnectTimeout    = flag.Duration("reconnect_timeout", 30*time.Second, "Reconnect timeout")
)

// http服务器配置
var (
	HttpTimeout = flag.Duration("http_timeout", 30*time.Second, "Http request timeout")
	GzipLevel   = flag.Int("gzip_level", 5, "GZip level")
)

const (
	BeginYear = 2025

	DefaultFilePermissions     = 0644
	DefaultDirectoryPermission = 0755

	LogName = "MAIN"

	EnvNoLogs              = "NO_LOGS"
	EnvAutoMigrate         = "AUTO_MIGRATE"
	EnvConfigFilePath      = "CONFIG_FILE_PATH"
	EnvHealthCheckInterval = "HEALTH_CHECK_INTERVAL"
	EnvHealthCheckTimeout  = "HEALTH_CHECK_TIMEOUT"
	EnvDeregisterAfter     = "DEREGISTER_AFTER"
	EnvServiceAddress      = "SERVICE_ADDRESS"
	EnvCenterAddress       = "CENTER_ADDRESS"
	EnvReconnectTimeout    = "RECONNECT_TIMEOUT"
	EnvEthName             = "ETH_NAME"
	EnvHttpTimeout         = "HTTP_TIMEOUT"
	EnvGzipLevel           = "GZIP_LEVEL"
)

func CheckFlags() {
	flag.Parse()
	utils.CheckBoolEnv(EnvNoLogs, NoLogs)
	utils.CheckBoolEnv(EnvAutoMigrate, AutoMigrate)
	utils.CheckStringEnv(EnvConfigFilePath, ConfigFilePath)
	utils.CheckStringEnv(EnvEthName, EthName)
	utils.CheckStringEnv(EnvHealthCheckInterval, HealthCheckInterval)
	utils.CheckStringEnv(EnvHealthCheckTimeout, HealthCheckTimeout)
	utils.CheckStringEnv(EnvDeregisterAfter, DeregisterAfter)
	*ServiceAddress = utils.GetLocalIP(*EthName)
	utils.CheckStringEnv(EnvServiceAddress, ServiceAddress)
	utils.CheckStringEnv(EnvCenterAddress, CenterAddress)
	utils.CheckDurationEnv(EnvReconnectTimeout, ReconnectTimeout)
	utils.CheckDurationEnv(EnvHttpTimeout, HttpTimeout)
	utils.CheckIntEnv(EnvGzipLevel, GzipLevel)
}
