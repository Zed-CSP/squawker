package models

import (
	"time"
)

type Message struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"userId" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	Volume    int       `json:"volume" db:"volume"`
	EchoCount int       `json:"echoCount" db:"echo_count"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
