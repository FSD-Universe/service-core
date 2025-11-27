package discovery

import (
	"time"
)

// 消息类型
const (
	MessageTypeDiscoveryRequest  = "DISCOVERY_REQ"
	MessageTypeDiscoveryResponse = "DISCOVERY_RESP"
	MessageTypeHeartbeat         = "HEARTBEAT"
	MessageTypeOnline            = "ONLINE"
	MessageTypeShutdown          = "SHUTDOWN"
)

var AllServices = "*"

// ServiceInfo 服务信息
type ServiceInfo struct {
	Name          string            `json:"name"`
	IP            string            `json:"ip"`
	Port          int               `json:"port"`
	Version       string            `json:"version"`
	LastHeartbeat time.Time         `json:"last_heartbeat"`
	Metadata      map[string]string `json:"metadata"`
}

// BroadcastMessage 广播消息
type BroadcastMessage struct {
	Type      string       `json:"type"`
	Sender    *ServiceInfo `json:"sender"`
	Timestamp int64        `json:"timestamp"`
	Target    string       `json:"target,omitempty"`
}
