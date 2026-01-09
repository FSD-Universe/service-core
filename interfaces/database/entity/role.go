// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import (
	"time"
)

type Role struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"size:64;not null"`
	Permission uint64 `gorm:"default:0;not null"`
	Comment    string `gorm:"type:text;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r *Role) GetId() uint {
	return r.ID
}

func (r *Role) SetId(id uint) {
	r.ID = id
}
