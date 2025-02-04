package api

import (
    "github.com/gin-gonic/gin"
    "squawker-backend/internal/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Enable CORS
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

    userController := controllers.NewUserController()
    messageController := controllers.NewMessageController()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // User routes
    router.POST("/register", userController.Register)
    router.POST("/login", userController.Login)

    // Message routes
    api := router.Group("/api")
    {
        api.GET("/messages", messageController.GetMessages)
        api.POST("/messages", messageController.CreateMessage)
    }
    
    return router
}
