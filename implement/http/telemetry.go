// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http && telemetry

// Package http
package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"half-nothing.cn/service-core/interfaces/config"
)

var SkipperHealthCheck = func(c echo.Context) bool {
	return c.Path() == "/health"
}

func SetTelemetry(e *echo.Echo, c *config.TelemetryConfig, skipper middleware.Skipper) {
	e.Use(otelecho.Middleware(c.Name, otelecho.WithSkipper(skipper)))
}
