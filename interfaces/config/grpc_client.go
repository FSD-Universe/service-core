// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

type GrpcClientConfig struct {
	EnableTLS  bool   `yaml:"enable_tls"`
	CA         string `yaml:"ca"`
	EnableMTLS bool   `yaml:"enable_mtls"`
	Cert       string `yaml:"cert"`
	Key        string `yaml:"key"`

	// 内部变量
	Credentials credentials.TransportCredentials
}

func (g *GrpcClientConfig) InitDefaults() {
	g.EnableTLS = false
	g.CA = ""
	g.EnableMTLS = false
	g.Cert = ""
	g.Key = ""
}

func (g *GrpcClientConfig) Verify() (bool, error) {
	if !g.EnableTLS {
		return true, nil
	}
	if g.CA == "" {
		return false, fmt.Errorf("ca is empty")
	}
	certPool := x509.NewCertPool()
	caCert, err := os.ReadFile(g.CA)
	if err != nil {
		return false, fmt.Errorf("read %s error: %v", g.CA, err)
	}
	if !certPool.AppendCertsFromPEM(caCert) {
		return false, fmt.Errorf("invalid ca")
	}
	if !g.EnableMTLS {
		g.Credentials = credentials.NewTLS(&tls.Config{RootCAs: certPool})
		return true, nil
	}
	if g.Cert == "" {
		return false, fmt.Errorf("cert is empty")
	}
	if g.Key == "" {
		return false, fmt.Errorf("key is empty")
	}
	cert, err := tls.LoadX509KeyPair(g.Cert, g.Key)
	if err != nil {
		return false, fmt.Errorf("load x509 key pair error: %v", err)
	}
	g.Credentials = credentials.NewTLS(&tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	})
	return true, nil
}
