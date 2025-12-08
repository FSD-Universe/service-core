//go:build database && controller

// Package entity
package entity

import "time"

type Controller struct {
	ID                  uint      `gorm:"primarykey"`
	UserId              uint      `gorm:"uniqueIndex:ui_user_id;not null"`
	InstructorId        uint      `gorm:"index;not null"`
	Guest               bool      `gorm:"default:false;not null"`
	UnderMonitor        bool      `gorm:"default:false;not null"`
	UnderSolo           bool      `gorm:"default:false;not null"`
	SoloUntil           time.Time `gorm:"default:null"`
	Tier2               bool      `gorm:"default:false;not null"`
	TotalControllerTime uint64    `gorm:"default:0;not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time

	// 外键定义
	User       *User       `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Instructor *Instructor `gorm:"foreignKey:InstructorId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (c *Controller) GetId() uint {
	return c.ID
}

func (c *Controller) SetId(id uint) {
	c.ID = id
}
