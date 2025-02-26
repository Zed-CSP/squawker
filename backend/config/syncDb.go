package config

import (
	"squawker-backend/internal/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.User{}, &models.Message{})
}
