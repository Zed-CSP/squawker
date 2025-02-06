package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"squawker-backend/config"
	"squawker-backend/internal/models"
)

func SignUp(c *gin.Context) {
	// Get username, password, and email from request body
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	// Validate request body
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password",})
		return
	}

	// Create user
	user := models.User{
		//Username: body.Username,
		Password: string(hash),
		Email:    body.Email,
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	//respond with success
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully",})
	
	
	
}