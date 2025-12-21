// Copyright (c) 2025 Half_nothing
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

// 服务发现广播配置
var (
	BroadcastPort     = flag.Int("broadcast_port", 9999, "Port for broadcast")
	HeartbeatInterval = flag.Duration("heartbeat_interval", 30*time.Second, "Heartbeat interval")
	ServiceTimeout    = flag.Duration("service_timeout", 90*time.Second, "Service timeout")
	CleanupInterval   = flag.Duration("cleanup_interval", 30*time.Second, "Cleanup interval")
	ReconnectTimeout  = flag.Duration("reconnect_timeout", 30*time.Second, "Reconnect timeout")
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

	EnvNoLogs            = "NO_LOGS"
	EnvAutoMigrate       = "AUTO_MIGRATE"
	EnvConfigFilePath    = "CONFIG_FILE_PATH"
	EnvBroadcastPort     = "BROADCAST_PORT"
	EnvHeartbeatInterval = "HEARTBEAT_INTERVAL"
	EnvServiceTimeout    = "SERVICE_TIMEOUT"
	EnvCleanupInterval   = "CLEANUP_INTERVAL"
	EnvReconnectTimeout  = "RECONNECT_TIMEOUT"
	EnvEthName           = "ETH_NAME"
	EnvHttpTimeout       = "HTTP_TIMEOUT"
	EnvGzipLevel         = "GZIP_LEVEL"
)

func CheckFlags() {
	flag.Parse()
	utils.CheckBoolEnv(EnvNoLogs, NoLogs)
	utils.CheckBoolEnv(EnvAutoMigrate, AutoMigrate)
	utils.CheckStringEnv(EnvConfigFilePath, ConfigFilePath)
	utils.CheckIntEnv(EnvBroadcastPort, BroadcastPort)
	utils.CheckDurationEnv(EnvHeartbeatInterval, HeartbeatInterval)
	utils.CheckDurationEnv(EnvServiceTimeout, ServiceTimeout)
	utils.CheckDurationEnv(EnvCleanupInterval, CleanupInterval)
	utils.CheckDurationEnv(EnvReconnectTimeout, ReconnectTimeout)
	utils.CheckStringEnv(EnvEthName, EthName)
	utils.CheckDurationEnv(EnvHttpTimeout, HttpTimeout)
	utils.CheckIntEnv(EnvGzipLevel, GzipLevel)
}
