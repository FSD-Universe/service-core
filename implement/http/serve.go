// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package http
package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/logger"
)

func Serve(lg logger.Interface, e *echo.Echo, c *config.HttpServerConfig) {
	protocol := "http"
	if c.SSLConfig.Enable {
		protocol = "https"
	}
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	lg.Infof("Starting %s server on %s", protocol, address)

	var err error
	if c.SSLConfig.Enable {
		err = e.StartTLS(
			address,
			c.SSLConfig.Cert,
			c.SSLConfig.Key,
		)
	} else {
		err = e.Start(address)
	}

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		lg.Fatalf("Http server error: %v", err)
	}
}
