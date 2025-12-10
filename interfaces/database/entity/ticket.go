//go:build database

// Package entity
package entity

import (
	"time"

	"half-nothing.cn/service-core/utils"
)

type Ticket struct {
	ID        uint    `gorm:"primarykey"`
	UserId    uint    `gorm:"index:idx_tickets_user_id;not null"`
	Type      int     `gorm:"default:0;not null"`
	Title     string  `gorm:"type:text;not null"`
	Content   string  `gorm:"type:text;not null"`
	Reply     *string `gorm:"type:text"`
	Replier   *uint   `gorm:"index:i_replier;default:null"`
	CreatedAt time.Time
	ClosedAt  time.Time

	// 外键定义
	User        *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	ReplierUser *User `gorm:"foreignKey:Replier;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (t *Ticket) GetId() uint {
	return t.ID
}

func (t *Ticket) SetId(id uint) {
	t.ID = id
}

type TicketType *utils.Enum[int, string]

var (
	TicketTypeFeature     = utils.NewEnum(0, "建议")
	TicketTypeBug         = utils.NewEnum(1, "bug")
	TicketTypeComplain    = utils.NewEnum(2, "投诉")
	TicketTypeRecognition = utils.NewEnum(3, "表扬")
	TicketTypeOtherType   = utils.NewEnum(4, "其他")
)

var TicketTypes = utils.NewEnums(
	TicketTypeFeature,
	TicketTypeBug,
	TicketTypeComplain,
	TicketTypeRecognition,
	TicketTypeOtherType,
)
