// Package utils
package utils

import (
	"net"
)

const (
	DefaultLocalIP     = "127.0.0.1"
	DefaultBroadcastIP = "255.255.255.255"
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

func GetBroadcastIP(ethName string, localIP string) string {
	ip := net.ParseIP(localIP)
	if ip == nil {
		return DefaultBroadcastIP
	}
	interfaces, err := net.Interfaces()
	if err != nil {
		return DefaultBroadcastIP
	}

	for _, iface := range interfaces {
		if iface.Name != ethName {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip4 := ipnet.IP.To4()
			if ip4 == nil {
				continue
			}

			if !ip4.Equal(ip) {
				continue
			}

			if ipnet.Mask == nil {
				continue
			}

			broadcast := make(net.IP, len(ip4))
			for i := range ip4 {
				broadcast[i] = ip4[i] | ^ipnet.Mask[i]
			}
			return broadcast.String()
		}
	}

	ip = ip.To4()
	if ip == nil {
		return DefaultBroadcastIP
	}
	ip[3] = 255
	return ip.String()
}
