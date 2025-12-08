//go:build database && user

// Package entity
package entity

import "time"

type User struct {
	ID             uint   `gorm:"primarykey"`
	Username       string `gorm:"uniqueIndex:ui_username;not null"`
	Email          string `gorm:"uniqueIndex:ui_email;not null"`
	Cid            int    `gorm:"uniqueIndex:ui_cid;not null"`
	Password       string `gorm:"size:128;not null"`
	ImageId        uint   `gorm:"default:null"`
	QQ             string `gorm:"size:16;default:null"`
	Rating         int    `gorm:"default:0;not null"`
	Permission     uint64 `gorm:"default:0;not null"`
	TotalPilotTime uint64 `gorm:"default:0;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastLogin      time.Time
	LastLoginIP    string `gorm:"size:128;default:null"`

	// 外键定义
	CurrentAvatar *Image      `gorm:"foreignKey:ID;references:ImageId"`
	FlightPlan    *FlightPlan `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (u *User) GetId() uint {
	return u.ID
}

func (u *User) SetId(id uint) {
	u.ID = id
}
