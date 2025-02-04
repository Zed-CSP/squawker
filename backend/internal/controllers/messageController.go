package controllers

import (
	"net/http"

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

func (ctrl *MessageController) GetMessages(c *gin.Context) {
	messages, err := ctrl.service.GetMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (ctrl *MessageController) CreateMessage(c *gin.Context) {
	// TODO: Implement message creation
	c.JSON(http.StatusOK, gin.H{"message": "Message creation endpoint"})
}
