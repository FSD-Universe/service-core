//go:build database

// Package entity
package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID             uint         `gorm:"primarykey"`
	Username       string       `gorm:"uniqueIndex:idx_users_username;size:64;not null"`
	Email          string       `gorm:"uniqueIndex:idx_users_email;size:128;not null"`
	Cid            uint         `gorm:"uniqueIndex:idx_users_cid;not null"`
	Password       string       `gorm:"size:128;not null"`
	ImageId        *uint        `gorm:"default:null"`
	QQ             *string      `gorm:"size:16;default:null"`
	Banned         bool         `gorm:"default:false;not null"`
	BannedUntil    sql.NullTime `gorm:"default:null"`
	Rating         int          `gorm:"default:0;not null"`
	Permission     uint64       `gorm:"default:0;not null"`
	TotalPilotTime uint64       `gorm:"default:0;not null"`
	LastLoginTime  sql.NullTime `gorm:"default:null"`
	LastLoginIP    *string      `gorm:"size:128;default:null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	// 外键定义
	CurrentAvatar *Image      `gorm:"foreignKey:ID;references:ImageId"`
	FlightPlan    *FlightPlan `gorm:"foreignKey:UserId;references:ID"`
	Roles         []*UserRole `gorm:"foreignKey:UserId;references:ID"`
}

func (u *User) GetId() uint {
	return u.ID
}

func (u *User) SetId(id uint) {
	u.ID = id
}
