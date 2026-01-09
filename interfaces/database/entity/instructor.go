// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import "time"

type Instructor struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint   `gorm:"uniqueIndex:idx_instructors_user_id;not null"`
	Email     string `gorm:"size:128;not null"`
	Twr       int    `gorm:"default:0;not null"`
	App       int    `gorm:"default:0;not null"`
	Ctr       int    `gorm:"default:0;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// 外键定义
	User        *User         `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Controllers []*Controller `gorm:"foreignKey:InstructorId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (i *Instructor) GetId() uint {
	return i.ID
}

func (i *Instructor) SetId(id uint) {
	i.ID = id
}
