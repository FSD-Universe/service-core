// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build telemetry

// Package database
package database

import (
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func ApplyDBTracing(db *gorm.DB, name string) error {
	return db.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithDBSystem(name)))
}
