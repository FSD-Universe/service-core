//go:build database

// Package entity
package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID               uint      `gorm:"primarykey"`
	PublisherId      uint      `gorm:"index:i_publisher_id;not null"`
	Type             int       `gorm:"default:0:not null"`
	Title            string    `gorm:"type:text;not null"`
	ImageId          uint      `gorm:"index:i_image_id;default:null"`
	ActiveTime       time.Time `gorm:"not null"`
	DepartureAirport string    `gorm:"size:64;not null"`
	ArrivalAirport   string    `gorm:"size:64;not null"`
	Route            string    `gorm:"type:text"`
	Distance         int       `gorm:"default:0"`
	Route2           string    `gorm:"type:text"`
	Distance2        int       `gorm:"default:0"`
	OpenFir          string    `gorm:"size:128;default:null"`
	Status           int       `gorm:"default:0;not null"`
	NOTAMS           string    `gorm:"type:text;not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt

	// 外键定义
	Publisher   *User                 `gorm:"foreignKey:PublisherId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Image       *Image                `gorm:"foreignKey:ImageId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Pilots      []*ActivityPilot      `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Facilities  []*ActivityFacility   `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Controllers []*ActivityController `gorm:"foreignKey:ActivityId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (a *Activity) GetId() uint {
	return a.ID
}

func (a *Activity) SetId(id uint) {
	a.ID = id
}
