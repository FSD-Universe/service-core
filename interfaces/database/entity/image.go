// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint   `gorm:"index:idx_images_user_id;not null"`
	Hashcode  string `gorm:"size:128;not null"`
	Filename  string `gorm:"size:128;not null"`
	Url       string `gorm:"type:text;not null"`
	Size      int64  `gorm:"default:0;not null"`
	MimeType  string `gorm:"size:128;not null"`
	Comment   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	// 外键定义
	User *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (i *Image) GetId() uint {
	return i.ID
}

func (i *Image) SetId(id uint) {
	i.ID = id
}
