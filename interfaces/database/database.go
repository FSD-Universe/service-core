// Package database
package database

import (
	"gorm.io/gorm"
	"half-nothing.cn/service-core/interfaces/database/entity"
)

func AutoMigrateAllTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Image{},
		&entity.History{},
		&entity.Role{},
		&entity.UserRole{},
		&entity.Announcement{},
		&entity.AuditLog{},
		&entity.FlightPlan{},
		&entity.ControllerApplication{},
		&entity.Instructor{},
		&entity.Controller{},
		&entity.ControllerRecord{},
		&entity.Ticket{},
		&entity.Activity{},
		&entity.ActivityPilot{},
		&entity.ActivityFacility{},
		&entity.ActivityController{},
		&entity.ActivityRecord{},
		&entity.ActivityCoordination{},
	)
}
