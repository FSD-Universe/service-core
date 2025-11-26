// Package global
package global

import (
	"flag"
	"time"
)

var (
	NoLogs         = flag.Bool("no_logs", false, "Disable logging to file")
	ConfigFilePath = flag.String("config", "./config.yaml", "Path to configuration file")
)

// 服务发现广播配置
var (
	BroadcastPort     = flag.Int("broadcast_port", 9999, "Port for broadcast")
	HeartbeatInterval = flag.Duration("heartbeat_interval", 30*time.Second, "Heartbeat interval")
	ServiceTimeout    = flag.Duration("service_timeout", 90*time.Second, "Service timeout")
	CleanupInterval   = flag.Duration("cleanup_interval", 30*time.Second, "Cleanup interval")
)

const (
	BeginYear = 2025

	DefaultFilePermissions     = 0644
	DefaultDirectoryPermission = 0755

	LogName = "MAIN"

	EnvNoLogs            = "NO_LOGS"
	EnvConfigFilePath    = "CONFIG_FILE_PATH"
	EnvBroadcastPort     = "BROADCAST_PORT"
	EnvHeartbeatInterval = "HEARTBEAT_INTERVAL"
	EnvServiceTimeout    = "SERVICE_TIMEOUT"
	EnvCleanupInterval   = "CLEANUP_INTERVAL"
)
