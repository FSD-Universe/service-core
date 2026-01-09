// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package cache
package cache

import "time"

type CachedItem[T any] struct {
	CachedData T
	ExpiredAt  time.Time
}

type Interface[K comparable, T any] interface {
	Set(key K, value T, expiredAt time.Time)
	SetWithTTL(key K, value T, ttl time.Duration)
	Get(key K) (T, bool)
	Del(key K)
	Close()
}
