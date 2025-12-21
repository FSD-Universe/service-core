// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build grpc

// Package config
package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

type GrpcTLSConfig struct {
	Enable     bool   `yaml:"enable"`
	EnableMTLS bool   `yaml:"enable_mtls"`
	CA         string `yaml:"ca"`
	Cert       string `yaml:"cert"`
	Key        string `yaml:"key"`

	// 内部变量
	Credentials credentials.TransportCredentials `yaml:"-"`
}

func (g *GrpcTLSConfig) InitDefaults() {
	g.Enable = false
	g.Cert = ""
	g.Key = ""
	g.EnableMTLS = false
	g.CA = ""
}

func (g *GrpcTLSConfig) Verify() (bool, error) {
	if !g.Enable {
		return true, nil
	}
	if g.Cert == "" {
		return false, fmt.Errorf("cert is empty")
	}
	if g.Key == "" {
		return false, fmt.Errorf("key is empty")
	}
	if !g.EnableMTLS {
		creds, err := credentials.NewServerTLSFromFile(g.Cert, g.Key)
		if err != nil {
			return false, fmt.Errorf("failed to create credentials: %v", err)
		}
		g.Credentials = creds
		return true, nil
	}
	if g.CA == "" {
		return false, fmt.Errorf("ca is empty")
	}
	serverCert, err := tls.LoadX509KeyPair(g.Cert, g.Key)
	if err != nil {
		return false, fmt.Errorf("failed to load server certificates: %v", err)
	}
	caCert, err := os.ReadFile(g.CA)
	if err != nil {
		return false, fmt.Errorf("read %s error: %v", g.CA, err)
	}
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return false, fmt.Errorf("failed to append ca certificates")
	}
	g.Credentials = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
		MinVersion:   tls.VersionTLS12,
	})
	return true, nil
}

type GrpcServerConfig struct {
	Host      string         `yaml:"host"`
	Port      int            `yaml:"port"`
	TLSConfig *GrpcTLSConfig `yaml:"tls"`
}

func (g *GrpcServerConfig) InitDefaults() {
	g.Host = "0.0.0.0"
	g.Port = 8081
	g.TLSConfig = &GrpcTLSConfig{}
	g.TLSConfig.InitDefaults()
}

func (g *GrpcServerConfig) Verify() (bool, error) {
	if g.Host == "" {
		return false, fmt.Errorf("host is empty")
	}
	if g.Port <= 0 {
		return false, fmt.Errorf("port must larger than 0")
	}
	if g.Port > 65535 {
		return false, fmt.Errorf("port must less than 65535")
	}
	if ok, err := g.TLSConfig.Verify(); !ok {
		return false, err
	}
	return true, nil
}
