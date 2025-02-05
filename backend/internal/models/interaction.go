package models

import (
	"time"
)

type Interaction struct {
	ID              int       `json:"id" db:"id"`
	UserID          int       `json:"userId" db:"user_id"`
	MessageID       int       `json:"messageId" db:"message_id"`
	InteractionType string    `json:"interactionType" db:"interaction_type"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
}
