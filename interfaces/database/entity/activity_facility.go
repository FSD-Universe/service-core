// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import "time"

type ActivityFacility struct {
	ID         uint   `gorm:"primarykey"`
	ActivityId uint   `gorm:"index:idx_activity_facilities_activity_id;not null"`
	MinRating  int    `gorm:"default:0;not null"`
	Callsign   string `gorm:"size:32;not null"`
	Frequency  string `gorm:"size:32;not null"`
	Tier2      bool   `gorm:"default:false;not null"`
	SortIndex  int    `gorm:"default:0;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// 外键定义
	Activity   *Activity           `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Controller *ActivityController `gorm:"foreignKey:FacilityId;references:ID"`
}

func (a *ActivityFacility) GetId() uint {
	return a.ID
}

func (a *ActivityFacility) SetId(id uint) {
	a.ID = id
}
