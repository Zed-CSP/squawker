package api

import (
    "github.com/gin-gonic/gin"
    "squawker-backend/internal/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    userController := controllers.NewUserController()
    messageController := controllers.NewMessageController()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // User routes
    router.POST("/register", userController.Register)
    router.POST("/login", userController.Login)

    // Message routes
    router.GET("/messages", messageController.GetMessages)
    router.POST("/messages", messageController.CreateMessage)
    
    return router
}
