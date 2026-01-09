// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package entity
package entity

import (
	"time"
)

type ActivityCoordination struct {
	ID               uint    `gorm:"primarykey"`
	ActivityId       uint    `gorm:"index:idx_activity_coordinations_activity_id;not null"`
	FacilityId       uint    `gorm:"index:idx_activity_coordinations_facility_id;not null"`
	LogonCode        *string `gorm:"size:8;default:null"`
	LogonNetwork     *string `gorm:"size:64;default:null"`
	PdcAvailable     bool    `gorm:"default:false;not null"`
	Transform        *string `gorm:"size:64;default:null"`
	TransformRemarks *string `gorm:"type:text;default:null"`
	Procedure        *string `gorm:"size:64;default:null"`
	Runway           *string `gorm:"size:64;default:null"`
	RunwayRemarks    *string `gorm:"type:text;default:null"`
	Altitude         *string `gorm:"size:64;default:null"`
	Remarks          *string `gorm:"type:text;default:null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	// 外键定义
	Activity *Activity         `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Facility *ActivityFacility `gorm:"foreignKey:FacilityId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (a *ActivityCoordination) GetId() uint {
	return a.ID
}

func (a *ActivityCoordination) SetId(id uint) {
	a.ID = id
}
