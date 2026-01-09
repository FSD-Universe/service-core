// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package database
package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"half-nothing.cn/service-core/interfaces/cleaner"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/database/entity"
	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/interfaces/logger"
)

func autoMigrateAllTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Image{},
		&entity.History{},
		&entity.Role{},
		&entity.UserRole{},
		&entity.Announcement{},
		&entity.AuditLog{},
		&entity.FlightPlan{},
		&entity.ControllerApplication{},
		&entity.Instructor{},
		&entity.Controller{},
		&entity.ControllerRecord{},
		&entity.Ticket{},
		&entity.Activity{},
		&entity.ActivityPilot{},
		&entity.ActivityFacility{},
		&entity.ActivityController{},
		&entity.ActivityRecord{},
		&entity.ActivityCoordination{},
	)
}

func InitDatabase(lg logger.Interface, c *config.DatabaseConfig) (cleaner.ShutdownCallback, *gorm.DB, error) {
	connection := c.DatabaseType.Data(lg, c)
	gormConfig := gorm.Config{}
	gormConfig.DefaultTransactionTimeout = c.QueryTimeoutDuration
	gormConfig.PrepareStmt = true
	gormConfig.TranslateError = true

	if lg.Level() <= slog.LevelDebug {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Error)
	} else {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	db, err := gorm.Open(connection, &gormConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("error occured while connecting to operation: %v", err)
	}

	if *global.AutoMigrate && autoMigrateAllTable(db) != nil {
		return nil, nil, fmt.Errorf("error occured while migrating database: %v", err)
	}

	dbPool, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("error occured while creating operation pool: %v", err)
	}

	dbPool.SetMaxIdleConns(c.MaxConnections / 5)
	dbPool.SetMaxOpenConns(c.MaxConnections)
	dbPool.SetConnMaxLifetime(c.ConnectionTimeoutDuration)

	err = dbPool.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("error occured while pinging operation: %v", err)
	}

	lg.Info("Database initialized and connection established")

	return func(ctx context.Context) error {
		lg.Infof("Closing database connection")
		timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		finish := make(chan struct{})
		go func() {
			defer close(finish)
			if err := dbPool.Close(); err != nil {
				lg.Errorf("Error occured while closing database connection: %v", err)
			}
		}()
		select {
		case <-timeoutCtx.Done():
			lg.Errorf("Timeout while closing database connection")
			return timeoutCtx.Err()
		case <-finish:
			return nil
		}
	}, db, nil
}
