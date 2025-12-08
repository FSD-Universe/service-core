//go:build database && activity

// Package entity
package entity

type ActivityController struct {
	ID         uint `gorm:"primarykey"`
	UserId     uint `gorm:"index:i_user_id;not null"`
	ActivityId uint `gorm:"index:i_activity_id;not null"`
	FacilityId uint `gorm:"index:i_facility_id;not null"`

	// 外键定义
	User     *User             `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Activity *Activity         `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Facility *ActivityFacility `gorm:"foreignKey:FacilityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (a *ActivityController) GetId() uint {
	return a.ID
}

func (a *ActivityController) SetId(id uint) {
	a.ID = id
}
