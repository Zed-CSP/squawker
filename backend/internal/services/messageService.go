package services

type MessageService struct{}

func NewMessageService() *MessageService {
    return &MessageService{}
}

// GetMessages returns a list of messages
func (s *MessageService) GetMessages() ([]map[string]interface{}, error) {
    // For now, return static messages
    // Later, this will fetch from the database
    messages := []map[string]interface{}{
        {
            "id":        1,
            "content":   "Hello from the backend!",
            "createdAt": "2024-02-04T12:00:00Z",
        },
        {
            "id":        2,
            "content":   "The API is working!",
            "createdAt": "2024-02-04T12:01:00Z",
        },
    }
    
    return messages, nil
}
