//go:build database

// Package entity
package entity

import "time"

type AuditLog struct {
	ID            uint          `gorm:"primarykey"`
	Subject       uint          `gorm:"index:i_subject;not null"`
	Object        string        `gorm:"type:text;not null"`
	Event         string        `gorm:"size:64;not null"`
	Ip            string        `gorm:"size:128;not null"`
	UserAgent     string        `gorm:"type:text;not null"`
	ChangeDetails *ChangeDetail `gorm:"type:text;serializer:json;default:null"`
	CreatedAt     time.Time

	// 外键定义
	User *User `gorm:"foreignKey:Subject;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

type ChangeDetail struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

func (a *AuditLog) GetId() uint {
	return a.ID
}

func (a *AuditLog) SetId(id uint) {
	a.ID = id
}
