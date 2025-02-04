package api

import (
    "github.com/gin-gonic/gin"
    "squawker-backend/internal/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    userController := controllers.NewUserController()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // User routes
    router.POST("/register", userController.Register)
    router.POST("/login", userController.Login)

    return router
}
