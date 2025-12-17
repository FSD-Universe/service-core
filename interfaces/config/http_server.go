// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http

// Package config
package config

import (
	"fmt"

	"half-nothing.cn/service-core/utils"
)

type HSTSConfig struct {
	Enable            bool `yaml:"enable"`
	MaxAge            int  `yaml:"max_age"`
	IncludeSubdomains bool `yaml:"include_subdomains"`
}

func (h *HSTSConfig) InitDefaults() {
	h.Enable = false
	h.MaxAge = 5184000
	h.IncludeSubdomains = false
}

func (h *HSTSConfig) Verify() (bool, error) {
	if !h.Enable {
		return true, nil
	}
	if h.MaxAge <= 0 {
		return false, fmt.Errorf("max_age must larger than 0")
	}
	return true, nil
}

type HttpTLSConfig struct {
	Enable     bool        `yaml:"enable"`
	Cert       string      `yaml:"cert"`
	Key        string      `yaml:"key"`
	ForceHttps bool        `yaml:"force_https"`
	HSTSConfig *HSTSConfig `yaml:"hsts"`
}

func (s *HttpTLSConfig) InitDefaults() {
	s.Enable = false
	s.Cert = ""
	s.Key = ""
	s.ForceHttps = false
	s.HSTSConfig = &HSTSConfig{}
	s.HSTSConfig.InitDefaults()
}

func (s *HttpTLSConfig) Verify() (bool, error) {
	if !s.Enable {
		return true, nil
	}
	if s.Cert == "" {
		return false, fmt.Errorf("cert is empty")
	}
	if s.Key == "" {
		return false, fmt.Errorf("key is empty")
	}
	if ok, err := s.HSTSConfig.Verify(); !ok {
		return false, err
	}
	return true, nil
}

type HttpServerConfig struct {
	Enable    bool           `yaml:"enable"`
	Host      string         `yaml:"host"`
	Port      int            `yaml:"port"`
	BodyLimit string         `yaml:"body_limit"`
	RateLimit int            `yaml:"rate_limit"`
	ProxyType int            `yaml:"proxy_type"`
	TrustIps  []string       `yaml:"trust_ips"`
	SSLConfig *HttpTLSConfig `yaml:"tls"`

	// 内部变量
	Type ProxyType `yaml:"-"`
}

func (h *HttpServerConfig) InitDefaults() {
	h.Enable = true
	h.Host = "0.0.0.0"
	h.Port = 8080
	h.BodyLimit = "5M"
	h.RateLimit = 20
	h.ProxyType = 0
	h.TrustIps = []string{"0.0.0.0/0"}
	h.SSLConfig = &HttpTLSConfig{}
	h.SSLConfig.InitDefaults()
}

func (h *HttpServerConfig) Verify() (bool, error) {
	if !h.Enable {
		return true, nil
	}
	if h.Host == "" {
		return false, fmt.Errorf("host is empty")
	}
	if h.Port <= 0 {
		return false, fmt.Errorf("port must larger than 0")
	}
	if h.Port > 65535 {
		return false, fmt.Errorf("port must less than 65535")
	}
	if !ProxyTypes.IsValidEnum(h.ProxyType) {
		return false, fmt.Errorf("proxy_type must be 0, 1 or 2")
	}
	h.Type = ProxyTypes.GetEnum(h.ProxyType)
	if h.RateLimit < 0 {
		return false, fmt.Errorf("rate_limit must be positive")
	}
	if ok, err := h.SSLConfig.Verify(); !ok {
		return false, err
	}
	return true, nil
}

type ProxyType *utils.Enum[int, string]

var (
	ProxyTypeDirect       = utils.NewEnum[int, string](0, "direct")
	ProxyTypeXFFHeader    = utils.NewEnum[int, string](1, "header")
	ProxyTypeRealIPHeader = utils.NewEnum[int, string](2, "real_ip_header")
)

var ProxyTypes = utils.NewEnums[int, string](
	ProxyTypeDirect,
	ProxyTypeXFFHeader,
	ProxyTypeRealIPHeader,
)
