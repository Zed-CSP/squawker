package api

import (
	"squawker-backend/internal/controllers"
	"squawker-backend/internal/middleware"
	"squawker-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Global middleware
	router.Use(middleware.ErrorHandler())

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	messageController := controllers.NewMessageController()
	userController := controllers.NewUserController()

	// Public routes with validation
	router.POST("/register",
		middleware.ValidateRequest(models.RegisterRequest{}),
		userController.Register,
	)
	router.POST("/login",
		middleware.ValidateRequest(models.LoginRequest{}),
		userController.Login,
	)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/messages", messageController.GetMessages)
		api.POST("/messages",
			middleware.ValidateRequest(models.CreateMessageRequest{}),
			messageController.CreateMessage,
		)
	}

	return router
}
