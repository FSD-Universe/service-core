//go:build database

// Package entity
package entity

import "time"

type Instructor struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint   `gorm:"uniqueIndex:ui_user_id;not null"`
	Email     string `gorm:"size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// 外键定义
	User        *User         `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Controllers []*Controller `gorm:"foreignKey:InstructorId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (i *Instructor) GetId() uint {
	return i.ID
}

func (i *Instructor) SetId(id uint) {
	i.ID = id
}
