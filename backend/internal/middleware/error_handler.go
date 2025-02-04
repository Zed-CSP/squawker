package middleware

import (
	"log"
	"net/http"
	"squawker-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrorHandler middleware handles errors globally
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			log.Printf("Error: %v", err)

			// Handle different types of errors
			switch err {
			case services.ErrDailyPostLimitExceeded:
				c.JSON(http.StatusTooManyRequests, ErrorResponse{
					Error:   "rate_limit_exceeded",
					Message: "You can only post once per day",
					Code:    429,
				})
			case services.ErrUserNotFound:
				c.JSON(http.StatusNotFound, ErrorResponse{
					Error:   "not_found",
					Message: "User not found",
					Code:    404,
				})
			case services.ErrInvalidCredentials:
				c.JSON(http.StatusUnauthorized, ErrorResponse{
					Error:   "invalid_credentials",
					Message: "Invalid username or password",
					Code:    401,
				})
			default:
				// Default to internal server error
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "internal_error",
					Message: "An unexpected error occurred",
					Code:    500,
				})
			}

			// Abort the request
			c.Abort()
		}
	}
}

// AuthMiddleware checks for valid authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Error:   "unauthorized",
				Message: "No authorization token provided",
				Code:    401,
			})
			c.Abort()
			return
		}

		// TODO: Implement proper JWT validation
		// For now, just check if token exists
		if len(token) < 7 || token[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Error:   "invalid_token",
				Message: "Invalid authorization token format",
				Code:    401,
			})
			c.Abort()
			return
		}

		// Set user ID in context (we'll implement proper JWT parsing later)
		c.Set("userID", 1) // Temporary hardcoded user ID
		c.Next()
	}
}
