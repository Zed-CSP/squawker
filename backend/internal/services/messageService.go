package services

import (
	"squawker-backend/internal/database"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type MessageService struct{}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) GetMessages() ([]Message, error) {
	query := `
		SELECT m.id, m.user_id, u.username, m.content, m.created_at
		FROM messages m
		JOIN users u ON m.user_id = u.id
		ORDER BY m.created_at DESC
		LIMIT 100
	`
	
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(
			&msg.ID,
			&msg.UserID,
			&msg.Username,
			&msg.Content,
			&msg.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (s *MessageService) CreateMessage(userID int, content string) error {
	// Check if user has already posted today
	var count int
	err := database.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM daily_posts 
		WHERE user_id = $1 
		AND post_date = CURRENT_DATE
	`).Scan(&count)
	
	if err != nil {
		return err
	}
	
	if count > 0 {
		return ErrDailyPostLimitExceeded
	}

	// Start transaction
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert message
	_, err = tx.Exec(`
		INSERT INTO messages (user_id, content)
		VALUES ($1, $2)
	`, userID, content)
	
	if err != nil {
		return err
	}

	// Update daily posts
	_, err = tx.Exec(`
		INSERT INTO daily_posts (user_id, post_date, post_count)
		VALUES ($1, CURRENT_DATE, 1)
		ON CONFLICT (user_id, post_date) 
		DO UPDATE SET post_count = daily_posts.post_count + 1
	`, userID)
	
	if err != nil {
		return err
	}

	return tx.Commit()
}
