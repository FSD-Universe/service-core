//go:build database

// Package entity
package entity

type Role struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"size:64;not null"`
	Permission uint64 `gorm:"default:0;not null"`
}

func (r *Role) GetId() uint {
	return r.ID
}

func (r *Role) SetId(id uint) {
	r.ID = id
}
