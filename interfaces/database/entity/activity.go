//go:build database

// Package entity
package entity

import (
	"time"

	"gorm.io/gorm"
	"half-nothing.cn/service-core/utils"
)

type Activity struct {
	ID               uint      `gorm:"primarykey"`
	PublisherId      uint      `gorm:"index:idx_activities_publisher_id;not null"`
	Type             int       `gorm:"default:0;not null"`
	Title            string    `gorm:"type:text;not null"`
	ImageId          *uint     `gorm:"index:idx_activities_image_id;default:null"`
	ActiveTime       time.Time `gorm:"not null"`
	DepartureAirport string    `gorm:"size:64;not null"`
	ArrivalAirport   string    `gorm:"size:64;not null"`
	Route            *string   `gorm:"type:text;default:null"`
	Distance         *int      `gorm:"default:null"`
	SecondRoute      *string   `gorm:"type:text;default:null"`
	SecondDistance   *int      `gorm:"default:null"`
	OpenFir          *string   `gorm:"size:128;default:null"`
	Status           int       `gorm:"default:0;not null"`
	NOTAMS           string    `gorm:"type:text;not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt

	// 外键定义
	Publisher   *User                 `gorm:"foreignKey:PublisherId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Image       *Image                `gorm:"foreignKey:ImageId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Pilots      []*ActivityPilot      `gorm:"foreignKey:ActivityId;references:ID"`
	Facilities  []*ActivityFacility   `gorm:"foreignKey:ActivityId;references:ID"`
	Controllers []*ActivityController `gorm:"foreignKey:ActivityId;references:ID"`
}

func (a *Activity) GetId() uint {
	return a.ID
}

func (a *Activity) SetId(id uint) {
	a.ID = id
}

type ActivityStatus *utils.Enum[int, string]

var (
	ActivityStatusRegistering ActivityStatus = utils.NewEnum(0, "报名中")
	ActivityStatusInTheEvent  ActivityStatus = utils.NewEnum(1, "活动中")
	ActivityStatusEnded       ActivityStatus = utils.NewEnum(2, "已结束")
)

var ActivityStatusManager = utils.NewEnums(
	ActivityStatusRegistering,
	ActivityStatusInTheEvent,
	ActivityStatusEnded,
)

type ActivityType *utils.Enum[int, string]

var (
	ActivityTypeOneWay     ActivityType = utils.NewEnum(0, "单向单站")
	ActivityTypeBothWay    ActivityType = utils.NewEnum(1, "双向双站")
	ActivityTypeFIROpenDay ActivityType = utils.NewEnum(2, "空域开放日")
)

var ActivityTypeManager = utils.NewEnums(
	ActivityTypeOneWay,
	ActivityTypeBothWay,
	ActivityTypeFIROpenDay,
)
