// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import "time"

type UserRole struct {
	ID        uint `gorm:"primarykey"`
	UserId    uint `gorm:"index:idx_user_roles_user_id;uniqueIndex:idx_user_roles_user_role;not null"`
	RoleId    uint `gorm:"index:idx_user_roles_role_id;uniqueIndex:idx_user_roles_user_role;not null"`
	CreatedAt time.Time

	// 外键定义
	User *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Role *Role `gorm:"foreignKey:RoleId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (u *UserRole) GetId() uint {
	return u.ID
}

func (u *UserRole) SetId(id uint) {
	u.ID = id
}
