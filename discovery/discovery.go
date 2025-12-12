// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package discovery
package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"half-nothing.cn/service-core/interfaces/discovery"
	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/interfaces/logger"
	"half-nothing.cn/service-core/utils"
)

// ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	logger        logger.Interface
	serviceName   string
	port          int
	broadcastPort int
	services      map[string]*discovery.ServiceInfo
	servicesMutex sync.RWMutex
	running       chan bool
	conn          *net.UDPConn
	broadcastIP   string
	localIP       string
	version       *utils.Version
	broadcastAddr *net.UDPAddr
}

// NewServiceDiscovery 创建新的服务发现实例
func NewServiceDiscovery(
	lg logger.Interface,
	serviceName string,
	port int,
	version *utils.Version,
) *ServiceDiscovery {
	return &ServiceDiscovery{
		logger:        logger.NewLoggerAdapter(lg, "service-discovery"),
		serviceName:   serviceName,
		port:          port,
		broadcastPort: *global.BroadcastPort,
		services:      make(map[string]*discovery.ServiceInfo),
		version:       version,
		running:       nil,
	}
}

// createMessage 创建广播消息
func (sd *ServiceDiscovery) createMessage(msgType string, target string) *discovery.BroadcastMessage {
	return &discovery.BroadcastMessage{
		Type:      msgType,
		Sender:    sd.getServiceInfo(),
		Timestamp: time.Now().Unix(),
		Target:    target,
	}
}

// getServiceInfo 获取当前服务信息
func (sd *ServiceDiscovery) getServiceInfo() *discovery.ServiceInfo {
	return &discovery.ServiceInfo{
		Name:          sd.serviceName,
		IP:            sd.localIP,
		Port:          sd.port,
		Version:       sd.version.String(),
		LastHeartbeat: time.Now(),
		Metadata:      make(map[string]string),
	}
}

// sendBroadcast 发送广播消息
func (sd *ServiceDiscovery) sendBroadcast(msg *discovery.BroadcastMessage) error {
	if sd.conn == nil {
		return fmt.Errorf("connection not initialized")
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = sd.conn.WriteToUDP(data, sd.broadcastAddr)
	return err
}

// Start 启动服务发现
func (sd *ServiceDiscovery) Start() error {
	// 初始化UDP连接
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", sd.broadcastPort))
	if err != nil {
		return err
	}

	sd.conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}

	sd.localIP = utils.GetLocalIP(*global.EthName)
	sd.broadcastIP = utils.GetBroadcastIP(*global.EthName, sd.localIP)
	sd.broadcastAddr, _ = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", sd.broadcastIP, sd.broadcastPort))
	sd.running = make(chan bool)

	// 启动监听协程
	go sd.listener()

	if err := sd.sendOnlineMessage(); err != nil {
		sd.logger.Errorf("Failed to send online message: %v", err)
		return err
	}

	// 启动心跳协程
	go sd.heartbeatBroadcaster()

	// 启动清理协程
	go sd.serviceCleanup()

	// 发送初始发现请求
	if err := sd.sendDiscoveryRequest(discovery.AllServices); err != nil {
		sd.logger.Errorf("Failed to send discovery request: %v", err)
		return err
	}

	sd.logger.Infof("Service discovery started: %s on %s:%d", sd.serviceName, sd.localIP, sd.port)

	return nil
}

// listener 监听广播消息
func (sd *ServiceDiscovery) listener() {
	buffer := make([]byte, 1024)

	for {
		select {
		case <-sd.running:
			return
		default:
			_ = sd.conn.SetReadDeadline(time.Now().Add(1 * time.Second))

			n, addr, err := sd.conn.ReadFromUDP(buffer)
			if err != nil {
				var netErr net.Error
				if errors.As(err, &netErr) && netErr.Timeout() {
					continue
				}
				sd.logger.Debugf("Read error: %v", err)
				continue
			}

			dataCopy := make([]byte, n)
			copy(dataCopy, buffer[:n])

			go sd.handleMessage(dataCopy, addr)
		}
	}
}

// handleMessage 处理接收到的消息
func (sd *ServiceDiscovery) handleMessage(data []byte, addr *net.UDPAddr) {
	// 创建数据副本以避免并发访问问题
	msg := &discovery.BroadcastMessage{}
	if err := json.Unmarshal(data, &msg); err != nil {
		sd.logger.Errorf("Failed to unmarshal message: %v", err)
		return
	}

	sd.logger.Debugf("Received message: %+v", msg)

	// 忽略自己的消息
	if msg.Sender.Name == sd.serviceName {
		sd.logger.Debugf("Ignoring own message: %+v", msg)
		return
	}

	switch msg.Type {
	case discovery.MessageTypeDiscoveryRequest:
		sd.handleDiscoveryRequest(msg, addr)
	case discovery.MessageTypeDiscoveryResponse:
		sd.registerService(msg.Sender)
	case discovery.MessageTypeOnline:
		sd.registerService(msg.Sender)
	case discovery.MessageTypeHeartbeat:
		sd.updateServiceHeartbeat(msg.Sender)
	case discovery.MessageTypeShutdown:
		sd.unregisterService(msg.Sender.Name)
	}
}

// QueryService 查询特定服务（阻塞式，等待响应）
func (sd *ServiceDiscovery) QueryService(serviceName string, timeout time.Duration) (*discovery.ServiceInfo, error) {
	// 先检查本地是否已有该服务
	if service := sd.GetService(serviceName); service != nil {
		return service, nil
	}

	// 发送查询请求
	if err := sd.sendDiscoveryRequest(serviceName); err != nil {
		return nil, err
	}

	// 等待服务注册
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if service := sd.GetService(serviceName); service != nil {
			return service, nil
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil, fmt.Errorf("service %s not found within timeout", serviceName)
}

// handleDiscoveryRequest 处理发现请求
func (sd *ServiceDiscovery) handleDiscoveryRequest(msg *discovery.BroadcastMessage, addr *net.UDPAddr) {
	// 如果查询的是本服务，或者查询所有服务，则回复
	if msg.Target == discovery.AllServices || msg.Target == sd.serviceName {
		response := sd.createMessage(discovery.MessageTypeDiscoveryResponse, "")
		data, err := json.Marshal(response)
		if err != nil {
			sd.logger.Errorf("Failed to marshal query response: %v", err)
			return
		}

		_, err = sd.conn.WriteToUDP(data, addr)
		if err != nil {
			sd.logger.Errorf("Failed to send query response: %v", err)
		}
	}
}

// sendDiscoveryRequest 发送发现请求
func (sd *ServiceDiscovery) sendDiscoveryRequest(target string) error {
	msg := sd.createMessage(discovery.MessageTypeDiscoveryRequest, target)
	if err := sd.sendBroadcast(msg); err != nil {
		sd.logger.Debugf("Failed to send discovery request: %v", err)
		return err
	}
	return nil
}

// registerService 注册服务
func (sd *ServiceDiscovery) registerService(serviceInfo *discovery.ServiceInfo) {
	sd.servicesMutex.Lock()
	defer sd.servicesMutex.Unlock()

	// 更新心跳时间
	serviceInfo.LastHeartbeat = time.Now()
	sd.services[serviceInfo.Name] = serviceInfo

	sd.logger.Infof("Service registered: %s(%s) at %s:%d", serviceInfo.Name, serviceInfo.Version, serviceInfo.IP, serviceInfo.Port)
}

// updateServiceHeartbeat 更新服务心跳
func (sd *ServiceDiscovery) updateServiceHeartbeat(serviceInfo *discovery.ServiceInfo) {
	sd.servicesMutex.Lock()
	defer sd.servicesMutex.Unlock()

	if existing, exists := sd.services[serviceInfo.Name]; exists {
		existing.LastHeartbeat = time.Now()
	} else {
		sd.registerService(serviceInfo)
	}
}

// unregisterService 注销服务
func (sd *ServiceDiscovery) unregisterService(serviceName string) {
	sd.servicesMutex.Lock()
	defer sd.servicesMutex.Unlock()

	delete(sd.services, serviceName)
	fmt.Printf("Service unregistered: %s", serviceName)
}

func (sd *ServiceDiscovery) sendOnlineMessage() error {
	msg := sd.createMessage(discovery.MessageTypeOnline, "")
	return sd.sendBroadcast(msg)
}

// heartbeatBroadcaster 心跳广播
func (sd *ServiceDiscovery) heartbeatBroadcaster() {
	ticker := time.NewTicker(*global.HeartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			msg := sd.createMessage(discovery.MessageTypeHeartbeat, "")
			if err := sd.sendBroadcast(msg); err != nil {
				fmt.Printf("Failed to send heartbeat: %v", err)
			}
		case <-sd.running:
			return
		}
	}
}

// serviceCleanup 服务清理
func (sd *ServiceDiscovery) serviceCleanup() {
	ticker := time.NewTicker(*global.CleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sd.servicesMutex.Lock()
			now := time.Now()
			for name, service := range sd.services {
				if now.Sub(service.LastHeartbeat) > *global.ServiceTimeout {
					delete(sd.services, name)
					fmt.Printf("Service expired: %s", name)
				}
			}
			sd.servicesMutex.Unlock()
		case <-sd.running:
			return
		}
	}
}

// GetService 获取服务信息
func (sd *ServiceDiscovery) GetService(serviceName string) *discovery.ServiceInfo {
	sd.servicesMutex.RLock()
	defer sd.servicesMutex.RUnlock()

	if service, exists := sd.services[serviceName]; exists {
		// 返回副本避免并发修改
		serviceCopy := *service
		return &serviceCopy
	}
	return nil
}

// GetAllServices 获取所有服务
func (sd *ServiceDiscovery) GetAllServices() []discovery.ServiceInfo {
	sd.servicesMutex.RLock()
	defer sd.servicesMutex.RUnlock()

	services := make([]discovery.ServiceInfo, 0, len(sd.services))
	for _, service := range sd.services {
		// 拷贝副本避免外部修改
		services = append(services, *service)
	}
	return services
}

// WaitForService 等待特定服务上线
func (sd *ServiceDiscovery) WaitForService(serviceName string, timeout time.Duration) (*discovery.ServiceInfo, error) {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		if service := sd.GetService(serviceName); service != nil {
			return service, nil
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil, fmt.Errorf("service %s not found within timeout", serviceName)
}

// Stop 停止服务发现
func (sd *ServiceDiscovery) Stop(ctx context.Context) error {
	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	done := make(chan struct{})

	go func() {
		close(sd.running)
		defer func(conn *net.UDPConn) {
			if conn != nil {
				_ = conn.Close()
			}
			close(done)
		}(sd.conn)

		// 发送下线通知
		msg := sd.createMessage(discovery.MessageTypeShutdown, "")
		err := sd.sendBroadcast(msg)

		if err != nil {
			sd.logger.Errorf("Failed to send shutdown message: %v", err)
			return
		}

		sd.logger.Infof("Service discovery stopped")
	}()

	select {
	case <-timeout.Done():
		return fmt.Errorf("timeout while stopping service discovery")
	case <-done:
		return nil
	}
}
