// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http

// Package http
package http

import (
	"log/slog"
	"net"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/logger"
)

// SetRealIPMethod 根据配置设置Echo服务器的IP提取方法
// 用于从HTTP请求中正确获取客户端真实IP地址
func SetRealIPMethod(lg logger.Interface, e *echo.Echo, c *config.HttpServerConfig) {
	switch c.Type {
	// 直连模式：直接从连接中获取IP地址
	case config.ProxyTypeDirect:
		e.IPExtractor = echo.ExtractIPDirect()

	// X-Forwarded-For头部模式：从X-Forwarded-For请求头中提取IP
	case config.ProxyTypeXFFHeader:
		trustOperations := make([]echo.TrustOption, 0, len(c.TrustIps))

		for _, ip := range c.TrustIps {
			_, network, err := net.ParseCIDR(ip)
			if err != nil {
				lg.Warnf("%s is not a valid CIDR string, skipping it", ip)
				continue
			}
			trustOperations = append(trustOperations, echo.TrustIPRange(network))
		}
		e.IPExtractor = echo.ExtractIPFromXFFHeader(trustOperations...)

	// Real-IP头部模式：从X-Real-IP请求头中提取IP
	case config.ProxyTypeRealIPHeader:
		e.IPExtractor = echo.ExtractIPFromRealIPHeader()

	// 默认情况：理论上不应该执行到这里
	default:
		lg.Error("If you see this message, shutdown the server immediately.This message should not be output in theory.")
		e.IPExtractor = echo.ExtractIPDirect()
	}
}

func SetEchoLogger(lg logger.Interface, e *echo.Echo) {
	loggerConfig := slogecho.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
	}
	e.Use(slogecho.NewWithConfig(lg.LogHandler(), loggerConfig))
}

func SetTimeoutConfig(e *echo.Echo, timeout time.Duration, skipper middleware.Skipper) {
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: timeout,
		Skipper: skipper,
	}))
}

func SetRecoverConfig(lg logger.Interface, e *echo.Echo) {
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(ctx echo.Context, err error, stack []byte) error {
			lg.Errorf("Recovered from a fatal error: %v, stack: %s", err, string(stack))
			return err
		},
	}))
}

func SetSecureConfig(e *echo.Echo, c *config.SSLConfig) {
	if c.ForceHttps {
		e.Use(middleware.HTTPSRedirect())
	}

	secureConfig := middleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
	}
	if c.HSTSConfig.Enable {
		secureConfig.HSTSMaxAge = c.HSTSConfig.MaxAge
		secureConfig.HSTSExcludeSubdomains = c.HSTSConfig.IncludeSubdomains
	}
	e.Use(middleware.SecureWithConfig(secureConfig))
}

func SetCORSConfig(e *echo.Echo) {
	e.Use(middleware.CORS())
}

func SetBodyLimitConfig(lg logger.Interface, e *echo.Echo, c *config.HttpServerConfig) {
	if c.BodyLimit != "" {
		e.Use(middleware.BodyLimit(c.BodyLimit))
	} else {
		lg.Warn("No body limit set, be aware of possible DDOS attacks")
	}
}

func SetGzipConfig(e *echo.Echo, level int, skipper middleware.Skipper) {
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:   level,
		Skipper: skipper,
	}))
}

func SetRateLimit(lg logger.Interface, e *echo.Echo, c *config.HttpServerConfig) {
	if c.RateLimit != 0 {
		ipPathLimiter := NewSlidingWindowLimiter(time.Minute, c.RateLimit)
		ipPathLimiter.StartCleanup(2 * time.Minute)
		e.Use(RateLimitMiddleware(ipPathLimiter, CombinedKeyFunc))
		lg.Infof("Rate limit: %d requests per minute", c.RateLimit)
	} else {
		lg.Warn("No rate limit was set, be aware of possible DDOS attacks")
	}
}

func SetEchoConfig(lg logger.Interface, e *echo.Echo, c *config.HttpServerConfig, httpTimeout time.Duration) {
	SetEchoLogger(lg, e)
	SetRealIPMethod(lg, e, c)
	SetTimeoutConfig(e, httpTimeout, nil)
	SetRecoverConfig(lg, e)
	SetSecureConfig(e, c.SSLConfig)
	SetCORSConfig(e)
	SetBodyLimitConfig(lg, e, c)
	SetGzipConfig(e, 5, nil)
	SetRateLimit(lg, e, c)
}
