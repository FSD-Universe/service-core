//go:build database

// Package entity
package entity

type ActivityRecord struct {
	ID         uint `gorm:"primarykey"`
	ActivityId int  `gorm:"index:i_activity_id;not null"`
	UserId     uint `gorm:"index:i_user_id;not null"`

	// 外键定义
	Activity *Activity `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	User     *User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (a *ActivityRecord) GetId() uint {
	return a.ID
}

func (a *ActivityRecord) SetId(id uint) {
	a.ID = id
}
