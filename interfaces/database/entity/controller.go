// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import (
	"database/sql"
	"time"
)

type Controller struct {
	ID                  uint          `gorm:"primarykey"`
	UserId              uint          `gorm:"uniqueIndex:idx_controllers_user_id;not null"`
	InstructorId        *uint         `gorm:"index;not null"`
	Guest               bool          `gorm:"default:false;not null"`
	UnderMonitor        bool          `gorm:"default:false;not null"`
	UnderSolo           bool          `gorm:"default:false;not null"`
	SoloUntil           *sql.NullTime `gorm:"default:null"`
	Tier2               bool          `gorm:"default:false;not null"`
	TotalControllerTime uint64        `gorm:"default:0;not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time

	// 外键定义
	User       *User       `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Instructor *Instructor `gorm:"foreignKey:InstructorId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (c *Controller) GetId() uint {
	return c.ID
}

func (c *Controller) SetId(id uint) {
	c.ID = id
}
