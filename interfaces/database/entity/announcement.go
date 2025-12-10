//go:build database

// Package entity
package entity

import (
	"time"

	"half-nothing.cn/service-core/utils"
)

type Announcement struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint   `gorm:"index:idx_announcements_user_id;not null"`
	Title     string `gorm:"type:text;not null"`
	Content   string `gorm:"type:text;not null"`
	Type      int    `gorm:"default:0;not null"`
	Important bool   `gorm:"default:false;not null"`
	ForceShow bool   `gorm:"default:false;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// 外键定义
	User *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (a *Announcement) GetId() uint {
	return a.ID
}

func (a *Announcement) SetId(id uint) {
	a.ID = id
}

type AnnouncementType *utils.Enum[int, string]

var (
	AnnouncementTypeNormal     AnnouncementType = utils.NewEnum(0, "普通公告")
	AnnouncementTypeController AnnouncementType = utils.NewEnum(1, "空管中心公告")
	AnnouncementTypeTechnical  AnnouncementType = utils.NewEnum(2, "技术组公告")
)

var AnnouncementTypeManager = utils.NewEnums(
	AnnouncementTypeNormal,
	AnnouncementTypeController,
	AnnouncementTypeTechnical,
)
