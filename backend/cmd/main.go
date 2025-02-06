package main

import (
	// "log"
	// "os"

	"squawker-backend/config"
	"squawker-backend/internal/controllers"
	"squawker-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
	config.SyncDB()
	// // Initialize database
	// db, err := database.InitDB()
	// if err != nil {
	// 	log.Fatal("Failed to initialize database:", err)
	// }
	// defer db.Close()

	// // Get port from Heroku environment
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
}

func main() {
	config.LoadEnv()
	config.ConnectToDB()

	r := gin.Default()

	// Initialize handlers
	wsHandler := handlers.NewWebSocketHandler()
	healthHandler := handlers.NewHealthHandler()

	// Routes
	r.GET("/ws", wsHandler.HandleWebSocket)
	r.GET("/health", healthHandler.CheckHealth)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)
	//r.POST("/login", controllers.Login)

	r.Run() // listen and serve on 0.0.0.0:8080
}
