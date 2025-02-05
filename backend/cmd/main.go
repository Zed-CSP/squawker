package main

import (
	"log"
	"os"
	initializers "squawker/backend/config"
	initializers ""
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
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

	// Setup and run gin server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
