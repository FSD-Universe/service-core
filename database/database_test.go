// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package database
package database

import (
	"context"
	"os"
	"testing"

	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/testutils"
)

func TestInitDatabase(t *testing.T) {
	lg := testutils.NewFakeLogger(t)
	lg.Infof("")
	c := &config.DatabaseConfig{}
	c.InitDefaults()
	if ok, err := c.Verify(); !ok {
		lg.Errorf("error occurred when handling database config, %v", err)
		return
	}
	*global.AutoMigrate = true
	shutdown, db, err := InitDatabase(lg, c)
	if err != nil {
		lg.Errorf("error occurred when initializing the database, %v", err)
		return
	}
	if db == nil {
		lg.Errorf("database is nil")
		return
	}
	_ = shutdown(context.Background())
	f, err := os.Stat(c.Database)
	if err != nil {
		lg.Errorf("error occurred when checking database file, %v", err)
		return
	}
	if !f.IsDir() {
		lg.Info("database file exists, deleting...")
		_ = os.Remove(c.Database)
	}
}
