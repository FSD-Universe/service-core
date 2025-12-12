// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import (
	"time"

	"half-nothing.cn/service-core/utils"
)

type ActivityPilot struct {
	ID         uint   `gorm:"primarykey"`
	ActivityId uint   `gorm:"index:idx_activity_pilots_activity_id;not null"`
	UserId     uint   `gorm:"index:idx_activity_pilots_user_id;not null"`
	Callsign   string `gorm:"size:64;not null"`
	Aircraft   string `gorm:"size:64;not null"`
	Status     int    `gorm:"default:0;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// 外键定义
	Activity *Activity `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	User     *User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (a *ActivityPilot) GetId() uint {
	return a.ID
}

func (a *ActivityPilot) SetId(id uint) {
	a.ID = id
}

type ActivityPilotStatus *utils.Enum[int, string]

var (
	ActivityPilotStatusSigned    ActivityPilotStatus = utils.NewEnum(0, "报名")
	ActivityPilotStatusClearance ActivityPilotStatus = utils.NewEnum(1, "放行")
	ActivityPilotStatusTakeoff   ActivityPilotStatus = utils.NewEnum(2, "起飞")
	ActivityPilotStatusLanding   ActivityPilotStatus = utils.NewEnum(3, "着陆")
)

var ActivityPilotManager = utils.NewEnums(
	ActivityPilotStatusSigned,
	ActivityPilotStatusClearance,
	ActivityPilotStatusTakeoff,
	ActivityPilotStatusLanding,
)
