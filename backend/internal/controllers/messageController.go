package controllers

import (
	"net/http"
	"squawker-backend/internal/models"
	"squawker-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	service *services.MessageService
}

func NewMessageController() *MessageController {
	return &MessageController{
		service: services.NewMessageService(),
	}
}

// GetMessages retrieves messages with pagination
func (ctrl *MessageController) GetMessages(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, _ := c.Get("userID")

	// Get pagination parameters
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")

	messages, err := ctrl.service.GetMessages(page, limit)
	if err != nil {
		c.Error(err)
		return
	}

	// Validate response
	var response []models.MessageResponse
	for _, msg := range messages {
		response = append(response, models.MessageResponse{
			ID:        msg.ID,
			Content:   msg.Content,
			Username:  msg.Username,
			CreatedAt: msg.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// CreateMessage creates a new message
func (ctrl *MessageController) CreateMessage(c *gin.Context) {
	// Get validated request from context
	req, exists := c.Get("validatedRequest")
	if !exists {
		c.Error(services.ErrInvalidRequest)
		return
	}

	createReq, ok := req.(models.CreateMessageRequest)
	if !ok {
		c.Error(services.ErrInvalidRequest)
		return
	}

	// Get user ID from context
	userID, _ := c.Get("userID")

	// Create message
	message, err := ctrl.service.CreateMessage(userID.(int), createReq.Content)
	if err != nil {
		c.Error(err)
		return
	}

	// Validate and send response
	response := models.MessageResponse{
		ID:        message.ID,
		Content:   message.Content,
		Username:  message.Username,
		CreatedAt: message.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
