//go:build database

// Package entity
package entity

import (
	"time"

	"gorm.io/gorm"
)

type ControllerRecord struct {
	ID           uint   `gorm:"primarykey"`
	UserId       uint   `gorm:"index:i_user_id;not null"`
	InstructorId uint   `gorm:"index:i_instructor_id;not null"`
	Type         int    `gorm:"default:0;not null"`
	Content      string `gorm:"type:text;not null"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt

	// 外键定义
	User       *User       `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Instructor *Instructor `gorm:"foreignKey:InstructorId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (c *ControllerRecord) GetId() uint {
	return c.ID
}

func (c *ControllerRecord) SetId(id uint) {
	c.ID = id
}
