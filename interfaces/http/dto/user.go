// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package dto
package dto

import (
	"fmt"

	"half-nothing.cn/service-core/interfaces/database/entity"
)

type BaseUserInfo struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Cid       uint   `json:"cid"`
	AvatarUrl string `json:"avatar_url"`
	QQ        string `json:"qq"`
}

func (b *BaseUserInfo) FromUserEntity(user *entity.User) *BaseUserInfo {
	b.Id = user.ID
	b.Username = user.Username
	b.Email = user.Email
	b.Cid = user.Cid
	if user.QQ != nil {
		b.QQ = *user.QQ
	}
	if user.CurrentAvatar != nil {
		b.AvatarUrl = user.CurrentAvatar.Url
	} else if b.QQ != "" {
		b.AvatarUrl = fmt.Sprintf("https://q2.qlogo.cn/headimg_dl?dst_uin=%s&spec=100", b.QQ)
	} else {
		b.AvatarUrl = ""
	}
	return b
}
