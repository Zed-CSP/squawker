package models

import "time"

// RegisterRequest represents the registration request body
type RegisterRequest struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,min=8,max=72,strongpassword"`
	Email    string `json:"email" validate:"required,email"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}

// CreateMessageRequest represents the create message request body
type CreateMessageRequest struct {
	Content string `json:"content" validate:"required,min=1,max=280,content"`
}

// Response models
type MessageResponse struct {
	ID        int       `json:"id" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

type AuthResponse struct {
	Token    string `json:"token" validate:"required"`
	Username string `json:"username" validate:"required"`
}
