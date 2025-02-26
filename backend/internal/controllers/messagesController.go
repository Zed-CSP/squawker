package controllers

import (
	"fmt"
	"net/http"

	"squawker-backend/config"
	"squawker-backend/internal/handlers"
	"squawker-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// WebSocketHandler reference
var WebSocketHandler *handlers.WebSocketHandler

// SetWebSocketHandler sets the WebSocketHandler for broadcasting
func SetWebSocketHandler(handler *handlers.WebSocketHandler) {
	WebSocketHandler = handler
}

func CreateMessage(c *gin.Context) {
	fmt.Println("Received message creation request")

	// Get the user ID from the context (set by the auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	// Get the message content from request body
	var body struct {
		Content string `json:"content" binding:"required"`
	}

	// Validate request body
	if err := c.Bind(&body); err != nil {
		fmt.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	fmt.Println("Message content:", body.Content)

	// Find the user to get their username
	var user models.User
	if result := config.DB.First(&user, userID); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Create message
	message := models.Message{
		Content:  body.Content,
		UserID:   user.ID,
		Username: user.Username,
	}

	result := config.DB.Create(&message)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create message",
			"details": result.Error.Error(),
		})
		return
	}

	// Broadcast the new message to all connected WebSocket clients
	if WebSocketHandler != nil {
		WebSocketHandler.BroadcastNewMessage(message)
	}

	// Return the created message
	c.JSON(http.StatusOK, gin.H{
		"message": "Message created successfully",
		"data": gin.H{
			"id":        message.ID,
			"content":   message.Content,
			"userId":    message.UserID,
			"username":  message.Username,
			"createdAt": message.CreatedAt,
		},
	})
}

func GetMessages(c *gin.Context) {
	var messages []models.Message

	result := config.DB.Order("created_at desc").Limit(50).Find(&messages)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
