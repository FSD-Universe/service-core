// Package global
package global

import (
	"flag"
	"time"

	"half-nothing.cn/service-core/utils"
)

var (
	Debug          = flag.Bool("debug", false, "Enable debug mode")
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

	EnvDebug             = "DEBUG"
	EnvNoLogs            = "NO_LOGS"
	EnvConfigFilePath    = "CONFIG_FILE_PATH"
	EnvBroadcastPort     = "BROADCAST_PORT"
	EnvHeartbeatInterval = "HEARTBEAT_INTERVAL"
	EnvServiceTimeout    = "SERVICE_TIMEOUT"
	EnvCleanupInterval   = "CLEANUP_INTERVAL"
	EnvEthName           = "ETH_NAME"
)

func CheckFlags() {
	flag.Parse()
	utils.CheckBoolEnv(EnvDebug, Debug)
	utils.CheckBoolEnv(EnvNoLogs, NoLogs)
	utils.CheckStringEnv(EnvConfigFilePath, ConfigFilePath)
	utils.CheckIntEnv(EnvBroadcastPort, BroadcastPort)
	utils.CheckDurationEnv(EnvHeartbeatInterval, HeartbeatInterval)
	utils.CheckDurationEnv(EnvServiceTimeout, ServiceTimeout)
	utils.CheckDurationEnv(EnvCleanupInterval, CleanupInterval)
	utils.CheckStringEnv(EnvEthName, EthName)
}
