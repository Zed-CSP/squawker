package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"squawker-backend/config"
	"squawker-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // ***************************** REMOVE IN PROD ****************************************
	},
}

// WebSocketMessage represents a message sent over WebSocket
type WebSocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// WebSocketHandler manages WebSocket connections
type WebSocketHandler struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan WebSocketMessage
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mutex      sync.Mutex
}

// NewWebSocketHandler creates a new WebSocketHandler
func NewWebSocketHandler() *WebSocketHandler {
	handler := &WebSocketHandler{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan WebSocketMessage),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}

	// Start the background goroutine
	go handler.run()
	return handler
}

// run processes WebSocket events in the background
func (h *WebSocketHandler) run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = true
			h.mutex.Unlock()

			// Send initial messages to the new client
			h.sendInitialMessages(client)

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
			}
			h.mutex.Unlock()

		case message := <-h.broadcast:
			h.mutex.Lock()
			for client := range h.clients {
				err := client.WriteJSON(message)
				if err != nil {
					log.Printf("Error broadcasting to client: %v", err)
					client.Close()
					delete(h.clients, client)
				}
			}
			h.mutex.Unlock()
		}
	}
}

// sendInitialMessages sends existing messages to a new client
func (h *WebSocketHandler) sendInitialMessages(client *websocket.Conn) {
	var messages []models.Message
	result := config.DB.Order("created_at desc").Limit(50).Find(&messages)
	if result.Error != nil {
		log.Printf("Error fetching initial messages: %v", result.Error)
		return
	}

	// Format the messages for JSON
	type MessageResponse struct {
		ID        uint      `json:"id"`
		Content   string    `json:"content"`
		UserID    uint      `json:"userId"`
		Username  string    `json:"username"`
		CreatedAt string    `json:"createdAt"` // Send as formatted string
	}

	formattedMessages := make([]MessageResponse, len(messages))
	for i, msg := range messages {
		formattedMessages[i] = MessageResponse{
			ID:        msg.ID,
			Content:   msg.Content,
			UserID:    msg.UserID,
			Username:  msg.Username,
			CreatedAt: msg.CreatedAt.Format(time.RFC3339), // ISO format
		}
	}

	client.WriteJSON(WebSocketMessage{
		Type:    "initial_messages",
		Payload: formattedMessages,
	})
}

// BroadcastNewMessage broadcasts a new message to all connected clients
func (h *WebSocketHandler) BroadcastNewMessage(message models.Message) {
	// Format the message for JSON
	formattedMessage := struct {
		ID        uint      `json:"id"`
		Content   string    `json:"content"`
		UserID    uint      `json:"userId"`
		Username  string    `json:"username"`
		CreatedAt string    `json:"createdAt"` // Send as formatted string
	}{
		ID:        message.ID,
		Content:   message.Content,
		UserID:    message.UserID,
		Username:  message.Username,
		CreatedAt: message.CreatedAt.Format(time.RFC3339), // ISO format
	}

	h.broadcast <- WebSocketMessage{
		Type:    "new_message",
		Payload: formattedMessage,
	}
}

// HandleWebSocket handles WebSocket connections
func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	// Register the new client
	h.register <- conn

	// Send initial connection message
	conn.WriteJSON(WebSocketMessage{
		Type:    "connection_established",
		Payload: "Connected to server",
	})

	// Handle incoming messages
	go func() {
		defer func() {
			h.unregister <- conn
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Error reading message: %v", err)
				}
				break
			}

			// Parse the message
			var wsMessage WebSocketMessage
			if err := json.Unmarshal(message, &wsMessage); err != nil {
				log.Printf("Error parsing message: %v", err)
				continue
			}

			// Handle different message types
			switch wsMessage.Type {
			case "ping":
				conn.WriteJSON(WebSocketMessage{
					Type:    "pong",
					Payload: nil,
				})
			}
		}
	}()
}
