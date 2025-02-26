package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content  string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	Username string `gorm:"not null"` // Store username for quick access
}
