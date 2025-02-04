package controllers

import (
	"net/http"
	"squawker-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// UserController struct to attach user methods
type UserController struct {
	service *services.UserService
}

// NewUserController creates a new controller
func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(),
	}
}

// Register handles user registration
func (ctrl *UserController) Register(c *gin.Context) {
	// TODO: Implement user registration
	c.JSON(http.StatusOK, gin.H{"message": "Registration endpoint"})
}

// Login handles user login
func (ctrl *UserController) Login(c *gin.Context) {
	// TODO: Implement user login
	c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
}
