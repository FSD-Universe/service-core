// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

func Any[T comparable](src []T, comparator func(element T) bool) bool {
	for _, v := range src {
		if comparator(v) {
			return true
		}
	}
	return false
}

func Find[T comparable](src []T, comparator func(element T) bool) T {
	for _, v := range src {
		if comparator(v) {
			return v
		}
	}
	var zero T
	return zero
}

func Filter[T comparable](src []T, filter func(element T) bool) []T {
	result := make([]T, 0, len(src))
	for _, v := range src {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T comparable, R comparable](src []T, mapper func(element T) R) []R {
	result := make([]R, len(src))
	for i := range src {
		result[i] = mapper(src[i])
	}
	return result
}

func ForEach[T comparable](src []T, callback func(index int, element T)) {
	for i, v := range src {
		callback(i, v)
	}
}

func ReverseForEach[T any](slice []T, f func(index int, value T)) {
	for i := len(slice) - 1; i >= 0; i-- {
		f(i, slice[i])
	}
}
