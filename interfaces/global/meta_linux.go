//go:build linux

package global

import "flag"

var EthName = flag.String("eth_name", "eth0", "Ethernet interface name")
