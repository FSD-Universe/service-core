// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package entity
package entity

import "time"

type ControllerApplicationTime struct {
	ID            uint      `gorm:"primarykey"`
	ApplicationId uint      `gorm:"index:idx_controller_application_time_application_id;not null"`
	Time          time.Time `gorm:"not null"`
	Selected      bool      `gorm:"default:false;not null"`

	// 外键定义
	Application *ControllerApplication `gorm:"foreignKey:ApplicationId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (c *ControllerApplicationTime) GetId() uint {
	return c.ID
}

func (c *ControllerApplicationTime) SetId(id uint) {
	c.ID = id
}
