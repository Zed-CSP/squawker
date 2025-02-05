package api

import (
	"database/sql"
	"squawker-backend/internal/controllers"
	"squawker-backend/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize controllers with db connection
	authController := controllers.NewAuthController(db)
	messageController := controllers.NewMessageController()

	// Public routes
	router.GET("/ws", HandleWebSocket)

	// Auth routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", middleware.AuthMiddleware(), authController.RefreshToken)
	}

	// Message routes - protected by auth middleware
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/messages", messageController.GetMessages)
		api.POST("/messages", messageController.CreateMessage)
	}

	// Add a health check endpoint
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return router
}
