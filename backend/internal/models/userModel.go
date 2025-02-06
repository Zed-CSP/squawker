package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//ID        uint   `gorm:"primaryKey"`
	//Username  string `gorm:"unique"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}