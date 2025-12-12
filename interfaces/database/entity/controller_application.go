// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import "time"

type ControllerApplication struct {
	ID           uint    `gorm:"primarykey"`
	UserId       uint    `gorm:"index:idx_controller_applications_user_id;not null"`
	Reason       string  `gorm:"type:text;not null"`
	Record       string  `gorm:"type:text;not null"`
	IsGuest      bool    `gorm:"default:false;not null"`
	Platform     string  `gorm:"size:16;not null"`
	ImageId      *uint   `gorm:"index:idx_controller_applications_image_id;default:null"`
	Status       int     `gorm:"default:0;not null"`
	Message      *string `gorm:"type:text;default:null"`
	InstructorId *uint   `gorm:"index:idx_controller_applications_instructor_id;default:null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// 外键定义
	User       *User                        `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Image      *Image                       `gorm:"foreignKey:ImageId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Instructor *Instructor                  `gorm:"foreignKey:InstructorId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Times      []*ControllerApplicationTime `gorm:"foreignKey:ApplicationId;references:ID"`
}

func (c *ControllerApplication) GetId() uint {
	return c.ID
}

func (c *ControllerApplication) SetId(id uint) {
	c.ID = id
}
