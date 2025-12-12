// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http && telemetry

// Package middleware
package middleware

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"half-nothing.cn/service-core/interfaces/config"
)

func SetTelemetry(e *echo.Echo, c *config.TelemetryConfig) {
	e.Use(otelecho.Middleware(c.Name))
}
