// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package config
package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"half-nothing.cn/service-core/interfaces/logger"
	"half-nothing.cn/service-core/utils"
)

type DatabaseType *utils.Enum[string, func(logger.Interface, *DatabaseConfig) gorm.Dialector]

var (
	MySQL      = utils.NewEnum("mysql", mySQLConnection)
	PostgreSQL = utils.NewEnum("postgresql", postgreSQLConnection)
	SQLite     = utils.NewEnum("sqlite3", sqliteConnection)
)

var databases = utils.NewEnums(MySQL, PostgreSQL, SQLite)

type DatabaseConfig struct {
	Type              string `yaml:"type"`
	Database          string `yaml:"database"`
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	EnableSSL         bool   `yaml:"enable_ssl"`
	ConnectionTimeout string `yaml:"connection_timeout"`
	QueryTimeout      string `yaml:"query_timeout"`
	MaxConnections    int    `yaml:"max_connections"`

	// 内部变量
	ConnectionTimeoutDuration time.Duration `yaml:"-"`
	QueryTimeoutDuration      time.Duration `yaml:"-"`
	DatabaseType              DatabaseType  `yaml:"-"`
}

func (d *DatabaseConfig) InitDefaults() {
	d.Type = "sqlite3"
	d.Database = "database.db"
	d.Host = ""
	d.Port = 0
	d.Username = ""
	d.Password = ""
	d.EnableSSL = false
	d.ConnectionTimeout = "1h"
	d.QueryTimeout = "10s"
	d.MaxConnections = 32
}

func (d *DatabaseConfig) Verify() (bool, error) {
	if !databases.IsValidEnum(d.Type) {
		return false, fmt.Errorf("database type %s is not allowed, support database is %v, please check the configuration file", d.Type, databases.GetEnums())
	}

	d.DatabaseType = databases.GetEnum(d.Type)

	if d.Port < 0 || d.Port > 65535 || (d.Port == 0 && d.DatabaseType != SQLite) {
		return false, fmt.Errorf("database port %d is invalid, please check the configuration file", d.Port)
	}

	if err := utils.ParseDuration(d.ConnectionTimeout, &d.ConnectionTimeoutDuration); err != nil {
		return false, fmt.Errorf("connection timeout %s is invalid, please check the configuration file", d.ConnectionTimeout)
	}

	if err := utils.ParseDuration(d.QueryTimeout, &d.QueryTimeoutDuration); err != nil {
		return false, fmt.Errorf("query timeout %s is invalid, please check the configuration file", d.QueryTimeout)
	}

	return true, nil
}

func mySQLConnection(lg logger.Interface, db *DatabaseConfig) gorm.Dialector {
	var enableSSL string
	if db.EnableSSL {
		enableSSL = "true"
	} else {
		enableSSL = "false"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&tls=%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Database,
		enableSSL,
	)
	lg.Debugf("Mysql Connection DSN %s", dsn)
	return mysql.Open(dsn)
}

func postgreSQLConnection(lg logger.Interface, db *DatabaseConfig) gorm.Dialector {
	var enableSSL string
	if db.EnableSSL {
		enableSSL = "enable"
	} else {
		enableSSL = "disable"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		db.Host,
		db.Username,
		db.Password,
		db.Database,
		db.Port,
		enableSSL,
	)
	lg.Debugf("PostgreSQL Connection DSN %s", dsn)
	return postgres.Open(dsn)
}

func sqliteConnection(_ logger.Interface, db *DatabaseConfig) gorm.Dialector {
	return sqlite.Open(db.Database)
}
