//go:build windows

package global

import "flag"

var EthName = flag.String("eth_name", "Ethernet", "Ethernet interface name")
