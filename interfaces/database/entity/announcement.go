//go:build database

// Package entity
package entity

import "time"

type Announcement struct {
	ID        uint   `gorm:"primarykey"`
	PublishId uint   `gorm:"index:i_publish_id;not null"`
	Title     string `gorm:"type:text:not null"`
	Content   string `gorm:"type:text:not null"`
	Type      int    `gorm:"default:0;not null"`
	Important bool   `gorm:"default:false;not null"`
	ForceShow bool   `gorm:"default:false;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// 外键定义
	User *User `gorm:"foreignKey:PublishId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (a *Announcement) GetId() uint {
	return a.ID
}

func (a *Announcement) SetId(id uint) {
	a.ID = id
}
