// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build linux || darwin

package global

import "flag"

var EthName = flag.String("eth_name", "eth0", "Ethernet interface name")
