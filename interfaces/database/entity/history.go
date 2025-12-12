// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import "time"

type History struct {
	ID         uint      `gorm:"primarykey"`
	UserId     uint      `gorm:"index:idx_histories_user_id;not null"`
	Callsign   string    `gorm:"size:16;not null"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    time.Time `gorm:"not null"`
	OnlineTime int       `gorm:"default:0;not null"`
	IsAtc      bool      `gorm:"default:0;not null"`
	CreatedAt  time.Time

	// 外键定义
	User *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (h *History) GetId() uint {
	return h.ID
}

func (h *History) SetId(id uint) {
	h.ID = id
}
