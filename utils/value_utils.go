// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "reflect"

// GetOrDefault 获取数据，如果数据为空则返回默认数据
func GetOrDefault[T any](data T, defaultData T) T {
	if reflect.ValueOf(data).IsZero() {
		return defaultData
	}
	return data
}
