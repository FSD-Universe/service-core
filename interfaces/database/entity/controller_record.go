// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import (
	"time"

	"gorm.io/gorm"
	"half-nothing.cn/service-core/utils"
)

type ControllerRecord struct {
	ID           uint   `gorm:"primarykey"`
	UserId       uint   `gorm:"index:idx_controller_records_user_id;not null"`
	InstructorId uint   `gorm:"index:idx_controller_records_instructor_id;not null"`
	Type         int    `gorm:"default:0;not null"`
	Content      string `gorm:"type:text;not null"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt

	// 外键定义
	User       *User       `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Instructor *Instructor `gorm:"foreignKey:InstructorId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (c *ControllerRecord) GetId() uint {
	return c.ID
}

func (c *ControllerRecord) SetId(id uint) {
	c.ID = id
}

type ControllerRecordType utils.EnumIntString

var (
	ControllerRecordInterview    ControllerRecordType = utils.NewEnum(0, "面试")
	ControllerRecordSimulator    ControllerRecordType = utils.NewEnum(1, "模拟机")
	ControllerRecordRatingChange ControllerRecordType = utils.NewEnum(2, "权限变动")
	ControllerRecordTraining     ControllerRecordType = utils.NewEnum(3, "训练内容")
	ControllerRecordUnderMonitor ControllerRecordType = utils.NewEnum(4, "UM权限变动")
	ControllerRecordSolo         ControllerRecordType = utils.NewEnum(5, "Solo权限变动")
	ControllerRecordGuest        ControllerRecordType = utils.NewEnum(6, "客座权限变动")
	ControllerRecordApplication  ControllerRecordType = utils.NewEnum(7, "管制员申请")
	ControllerRecordOther        ControllerRecordType = utils.NewEnum(8, "其他未定义内容")
)

var ControllerRecordManager = utils.NewEnums(
	ControllerRecordInterview,
	ControllerRecordSimulator,
	ControllerRecordRatingChange,
	ControllerRecordTraining,
	ControllerRecordUnderMonitor,
	ControllerRecordSolo,
	ControllerRecordGuest,
	ControllerRecordApplication,
	ControllerRecordOther,
)
