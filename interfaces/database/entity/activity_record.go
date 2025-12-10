//go:build database

// Package entity
package entity

import "time"

type ActivityRecord struct {
	ID         uint `gorm:"primarykey"`
	ActivityId int  `gorm:"index:idx_activity_records_activity_id;not null"`
	UserId     uint `gorm:"index:idx_activity_records_user_id;not null"`
	CreatedAt  time.Time

	// 外键定义
	Activity *Activity `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	User     *User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (a *ActivityRecord) GetId() uint {
	return a.ID
}

func (a *ActivityRecord) SetId(id uint) {
	a.ID = id
}
