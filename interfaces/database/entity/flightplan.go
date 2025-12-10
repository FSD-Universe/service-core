//go:build database

// Package entity
package entity

import (
	"fmt"
	"time"
)

type FlightPlan struct {
	ID                uint   `gorm:"primarykey"`
	UserId            uint   `gorm:"uniqueIndex:idx_flight_plans_user_id;not null"`
	Callsign          string `gorm:"size:16;not null"`
	FlightRule        string `gorm:"size:4;not null"`
	Aircraft          string `gorm:"size:128;not null"`
	Tas               int    `gorm:"default:0;not null"`
	DepartureAirport  string `gorm:"size:8;not null"`
	PlanDepartureTime string `gorm:"size:8;not null"`
	AtcDepartureTime  string `gorm:"size:8;not null"`
	CruiseAltitude    string `gorm:"size:8;not null"`
	ArrivalAirport    string `gorm:"size:8;not null"`
	RouteTimeHour     string `gorm:"size:2;not null"`
	RouteTimeMinute   string `gorm:"size:2;not null"`
	AirTimeHour       string `gorm:"size:2;not null"`
	AirTimeMinute     string `gorm:"size:2;not null"`
	AlternateAirport  string `gorm:"size:8;not null"`
	Remarks           string `gorm:"type:text;not null"`
	Route             string `gorm:"type:text;not null"`
	Locked            bool   `gorm:"default:false;not null"`
	FromWeb           bool   `gorm:"default:false;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	// 外键定义
	User *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (f *FlightPlan) GetId() uint {
	return f.ID
}

func (f *FlightPlan) SetId(id uint) {
	f.ID = id
}

func (f *FlightPlan) ToRawMessage() string {
	return fmt.Sprintf("")
}
