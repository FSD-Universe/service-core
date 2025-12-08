//go:build database

// Package entity
package entity

import "time"

type History struct {
	ID         uint      `gorm:"primarykey"`
	UserId     uint      `gorm:"index:i_user_id;not null"`
	Callsign   string    `gorm:"size:16;not null"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    time.Time `gorm:"not null"`
	OnlineTime int       `gorm:"default:0;not null"`
	IsAtc      bool      `gorm:"default:0;not null"`
	CreatedAt  time.Time

	// 外键定义
	User *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (h *History) GetId() uint {
	return h.ID
}

func (h *History) SetId(id uint) {
	h.ID = id
}
