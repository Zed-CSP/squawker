package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	// "squawker-backend/internal/models"
)

func RequireAuth(c *gin.Context) {
	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")

	// Check if the header exists and has the correct format
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// Extract the token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
		return
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		// Get the user ID from the token
		userID := uint(claims["sub"].(float64))

		// Set the user ID in the context
		c.Set("userID", userID)

		// Continue to the next handler
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
}
