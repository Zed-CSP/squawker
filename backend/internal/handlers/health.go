package handlers

import (
	"squawker-backend/config"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) CheckHealth(c *gin.Context) {
	// Get the underlying *sql.DB from GORM
	sqlDB, err := config.DB.DB()
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Could not get database instance",
			"error":   err.Error(),
		})
		return
	}

	// Ping the database
	err = sqlDB.Ping()
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Database connection failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":   "healthy",
		"database": "connected",
	})
}
