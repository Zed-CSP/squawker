package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "squawker-backend/internal/services"
)

// UserController struct to attach user methods
type UserController struct{}

// NewUserController creates a new controller
func NewUserController() *UserController {
    return &UserController{}
}

// Register handles user registration
func (ctrl UserController) Register(c *gin.Context) {
    // Call to service layer
    // Assume registration logic is performed and responds accordingly
    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Login handles user login
func (ctrl UserController) Login(c *gin.Context) {
    // Assume login logic is performed and responds accordingly
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
