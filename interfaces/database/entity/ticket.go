//go:build database

// Package entity
package entity

import "time"

type Ticket struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint   `gorm:"index:i_user_id;not null"`
	Type      int    `gorm:"default:0;not null"`
	Title     string `gorm:"type:text;not null"`
	Content   string `gorm:"type:text;not null"`
	Reply     string `gorm:"type:text;not null"`
	Replier   uint   `gorm:"index:i_replier;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  time.Time

	// 外键定义
	User        *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	ReplierUser *User `gorm:"foreignKey:Replier;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (t *Ticket) GetId() uint {
	return t.ID
}

func (t *Ticket) SetId(id uint) {
	t.ID = id
}
