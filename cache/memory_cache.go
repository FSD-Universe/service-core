// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package cache
package cache

import (
	"sync"
	"time"

	"half-nothing.cn/service-core/interfaces/cache"
	"half-nothing.cn/service-core/utils"
)

func isOutDate[T any](data *cache.CachedItem[T]) bool {
	return data.ExpiredAt.Before(time.Now())
}

type MemoryCache[K comparable, T any] struct {
	cacheMap map[K]*cache.CachedItem[T]
	cleaner  *utils.IntervalActuator
	lock     sync.RWMutex
}

func NewMemoryCache[K comparable, T any](cleanInterval time.Duration) *MemoryCache[K, T] {
	if cleanInterval <= 0 {
		cleanInterval = 30 * time.Minute
	}
	cached := &MemoryCache[K, T]{
		cacheMap: make(map[K]*cache.CachedItem[T]),
		lock:     sync.RWMutex{},
	}
	cached.cleaner = utils.NewIntervalActuator(cleanInterval, cached.CleanExpiredData)
	return cached
}

func (memoryCache *MemoryCache[K, T]) CleanExpiredData() {
	memoryCache.lock.Lock()
	defer memoryCache.lock.Unlock()

	for key, value := range memoryCache.cacheMap {
		if isOutDate(value) {
			delete(memoryCache.cacheMap, key)
		}
	}
}

func (memoryCache *MemoryCache[K, T]) Set(key K, value T, expiredAt time.Time) {
	if expiredAt.Before(time.Now()) {
		return
	}
	var zero K
	if key == zero {
		return
	}
	memoryCache.lock.Lock()
	memoryCache.cacheMap[key] = &cache.CachedItem[T]{CachedData: value, ExpiredAt: expiredAt}
	memoryCache.lock.Unlock()
}

func (memoryCache *MemoryCache[K, T]) SetWithTTL(key K, value T, ttl time.Duration) {
	expiredAt := time.Now().Add(ttl)
	memoryCache.Set(key, value, expiredAt)
}

func (memoryCache *MemoryCache[K, T]) Get(key K) (T, bool) {
	var zero K
	if key == zero {
		var zero T
		return zero, false
	}
	memoryCache.lock.RLock()
	defer memoryCache.lock.RUnlock()
	val, ok := memoryCache.cacheMap[key]
	if ok && isOutDate(val) {
		var zero T
		return zero, false
	}
	if val == nil {
		var zero T
		return zero, false
	}
	return val.CachedData, ok
}

func (memoryCache *MemoryCache[K, T]) Del(key K) {
	var zero K
	if key == zero {
		return
	}
	memoryCache.lock.Lock()
	delete(memoryCache.cacheMap, key)
	memoryCache.lock.Unlock()
}

func (memoryCache *MemoryCache[K, T]) Close() {
	memoryCache.cleaner.Stop()
}
