//go:build database

// Package entity
package entity

type ActivityPilot struct {
	ID         uint   `gorm:"primarykey"`
	ActivityId uint   `gorm:"index:i_activity_id;not null"`
	UserId     uint   `gorm:"index:i_user_id;not null"`
	Callsign   string `gorm:"size:64;not null"`
	Aircraft   string `gorm:"size:64;not null"`
	Status     int    `gorm:"default:0;not null"`

	// 外键定义
	Activity *Activity `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	User     *User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (a *ActivityPilot) GetId() uint {
	return a.ID
}

func (a *ActivityPilot) SetId(id uint) {
	a.ID = id
}
