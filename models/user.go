package models

import (
	"time"
)

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string `json:"name"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	RoleID    *uint  `json:"role_id"  gorm:"null"`
	Role      Role
}

type Role struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type RolePermission struct {
	RoleID       int `gorm:"primaryKey"`
	PermissionID int `gorm:"primaryKey"`
}
