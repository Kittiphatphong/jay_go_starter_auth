package models

import (
	"time"
)

type Partner struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string `json:"name"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Status    bool   `json:"status" gorm:"default:true"`
}
