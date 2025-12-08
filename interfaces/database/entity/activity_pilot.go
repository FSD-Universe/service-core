//go:build database && activity

// Package entity
package entity

import "half-nothing.cn/service-core/utils"

type ActivityPilot struct {
	ID         uint   `gorm:"primarykey"`
	ActivityId uint   `gorm:"index:i_activity_id;not null"`
	UserId     uint   `gorm:"index:i_user_id;not null"`
	Callsign   string `gorm:"size:64;not null"`
	Aircraft   string `gorm:"size:64;not null"`
	Status     int    `gorm:"default:0;not null"`

	// 外键定义
	Activity *Activity `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	User     *User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
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
