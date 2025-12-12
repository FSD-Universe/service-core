// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build windows

package global

import "flag"

var EthName = flag.String("eth_name", "Ethernet", "Ethernet interface name")
