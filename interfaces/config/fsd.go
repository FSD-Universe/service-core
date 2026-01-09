// Copyright (c) 2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build fsd

// Package config
package config

import (
	"fmt"
	"strings"
	"text/template"
	"time"

	"half-nothing.cn/service-core/utils"
)

type FsdRangeLimit struct {
	RefuseOutRange       bool `yaml:"refuse_out_range"`
	Observer             int  `json:"observer"`
	Delivery             int  `json:"delivery"`
	Ground               int  `json:"ground"`
	Tower                int  `json:"tower"`
	Approach             int  `json:"approach"`
	Center               int  `json:"center"`
	Apron                int  `json:"apron"`
	Supervisor           int  `json:"supervisor"`
	Administrator        int  `json:"administrator"`
	FlightServiceStation int  `json:"flight_service_station"`
}

func (f *FsdRangeLimit) InitDefaults() {
	f.RefuseOutRange = false
	f.Observer = 300
	f.Delivery = 20
	f.Ground = 20
	f.Tower = 50
	f.Approach = 150
	f.Center = 600
	f.Apron = 20
	f.Supervisor = 300
	f.Administrator = 300
	f.FlightServiceStation = 1500
}

func (f *FsdRangeLimit) Verify() (bool, error) {
	return true, nil
}

type FsdConfig struct {
	FsdName               string         `yaml:"fsd_name"`
	Host                  string         `yaml:"host"`
	Port                  int            `yaml:"port"`
	AirportDataFile       string         `yaml:"airport_data_file"`
	HeartBeatInterval     string         `yaml:"heartbeat_interval"`
	WhazzupExpireInterval string         `yaml:"whazzup_expire_interval"`
	SessionExpireInterval string         `yaml:"session_expire_interval"`
	MaxClient             int            `yaml:"max_client"`
	VisualRangeLimit      *FsdRangeLimit `yaml:"visual_range_limit"`
	MotdTemplate          []string       `yaml:"motd"`

	// 内部变量
	HeartBeatDuration     time.Duration `yaml:"-"`
	WhazzupExpireDuration time.Duration `yaml:"-"`
	SessionExpireDuration time.Duration `yaml:"-"`
	Motd                  []byte        `yaml:"-"`
}

func (f *FsdConfig) InitDefaults() {
	f.FsdName = "SkyMesh-FSD"
	f.Host = "0.0.0.0"
	f.Port = 6809
	f.AirportDataFile = "data/airports.json"
	f.HeartBeatInterval = "60s"
	f.WhazzupExpireInterval = "15s"
	f.SessionExpireInterval = "40s"
	f.MaxClient = 128
	f.VisualRangeLimit = &FsdRangeLimit{}
	f.VisualRangeLimit.InitDefaults()
	f.MotdTemplate = []string{
		"Welcome to use {{.Name}} v{{.Version}}",
	}
}

func (f *FsdConfig) Verify() (bool, error) {
	if f.Host == "" {
		return false, fmt.Errorf("host is empty")
	}
	if f.Port <= 0 {
		return false, fmt.Errorf("port must larger than 0")
	}
	if f.Port > 65535 {
		return false, fmt.Errorf("port must less than 65535")
	}
	if f.MaxClient <= 0 {
		return false, fmt.Errorf("max_client must larger than 0")
	}
	if err := utils.ParseDuration(f.HeartBeatInterval, &f.HeartBeatDuration); err != nil {
		return false, fmt.Errorf("heartbeat interval %s is invalid, please check the configuration file", f.HeartBeatInterval)
	}
	if err := utils.ParseDuration(f.WhazzupExpireInterval, &f.WhazzupExpireDuration); err != nil {
		return false, fmt.Errorf("whazzup expire interval %s is invalid, please check the configuration file", f.WhazzupExpireInterval)
	}
	if err := utils.ParseDuration(f.SessionExpireInterval, &f.SessionExpireDuration); err != nil {
		return false, fmt.Errorf("session expire interval %s is invalid, please check the configuration file", f.SessionExpireInterval)
	}
	if ok, err := f.VisualRangeLimit.Verify(); !ok {
		return false, err
	}
	return true, nil
}

func (f *FsdConfig) InitMotd(version string) error {
	motds := make([]string, len(f.MotdTemplate))
	data := struct {
		Name    string
		Version string
	}{
		Name:    f.FsdName,
		Version: version,
	}
	for i, c := range f.MotdTemplate {
		temp, err := template.New(fmt.Sprintf("motd%d", i)).Parse(c)
		if err != nil {
			return err
		}
		var sb strings.Builder
		if err := temp.Execute(&sb, data); err != nil {
			return err
		}
		motds[i] = sb.String()
	}
	f.Motd = []byte(strings.Join(motds, "\n"))
	return nil
}
