// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import (
	"net"
)

const (
	DefaultLocalIP = "127.0.0.1"
)

func GetLocalIP(ethName string) string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return DefaultLocalIP
	}

	for _, iface := range interfaces {
		if iface.Name != ethName {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		extractIp := func(addr net.Addr) string {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				return ""
			}

			if ipnet.IP.IsLoopback() {
				return ""
			}

			if ip := ipnet.IP.To4(); ip != nil {
				return ip.String()
			}

			return ""
		}

		for _, addr := range addrs {
			if ip := extractIp(addr); ip != "" {
				return ip
			}
		}
	}

	return DefaultLocalIP
}
