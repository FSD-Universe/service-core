//go:build database

// Package entity
package entity

type UserRole struct {
	ID     uint `gorm:"primarykey"`
	UserId uint `gorm:"index:i_user_id;not null"`
	RoleId uint `gorm:"index:i_role_id;not null"`

	// 外键定义
	User *User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
	Role *Role `gorm:"foreignKey:RoleId;references:ID;constraint:OnUpdate:cascade,OnDelete:cascade"`
}

func (u *UserRole) GetId() uint {
	return u.ID
}

func (u *UserRole) SetId(id uint) {
	u.ID = id
}
