// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "reflect"

// GetPointerDataOrDefault 获取指针的值，如果指针为空则返回默认值
func GetPointerDataOrDefault[T any](data *T, defaultData T) T {
	if reflect.ValueOf(data).IsZero() {
		return defaultData
	}
	return *data
}

// GetPointerData 获取指针的值, 如果指针为空则返回零值
func GetPointerData[T any](data *T) T {
	val := reflect.ValueOf(data)
	if val.IsZero() {
		return reflect.Zero(val.Type().Elem()).Interface().(T)
	}
	return *data
}
