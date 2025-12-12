// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import "time"

type ActivityController struct {
	ID         uint `gorm:"primarykey"`
	UserId     uint `gorm:"index:idx_activity_controllers_user_id;not null"`
	ActivityId uint `gorm:"index:idx_activity_controllers_activity_id;not null"`
	FacilityId uint `gorm:"index:idx_activity_controllers_facility_id;not null"`
	CreatedAt  time.Time

	// 外键定义
	User     *User             `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Activity *Activity         `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Facility *ActivityFacility `gorm:"foreignKey:FacilityId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (a *ActivityController) GetId() uint {
	return a.ID
}

func (a *ActivityController) SetId(id uint) {
	a.ID = id
}
