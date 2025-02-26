package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"squawker-backend/config"
	"squawker-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Log the request
	fmt.Println("Received signup request at:", c.Request.URL.Path)

	// Get username, password, and email from request body
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	// Validate request body
	if err := c.Bind(&body); err != nil {
		fmt.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	fmt.Println("Request body:", body.Username, body.Email)

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := models.User{
		Username: body.Username,
		Password: string(hash),
		Email:    body.Email,
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": result.Error.Error()})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week
	})

	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the user and token
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"token": tokenString,
	})
}

func Login(c *gin.Context) {
	// Get username and password from request body
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Find the user
	var user models.User
	result := config.DB.Where("username = ?", body.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week
	})

	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the user and token
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"token": tokenString,
	})
}

func VerifyToken(c *gin.Context) {
	// The middleware already verified the token
	// Return the user ID from the context
	userID, _ := c.Get("userID")

	// Find the user
	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
