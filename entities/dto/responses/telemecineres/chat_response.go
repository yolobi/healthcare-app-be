package telemecineres

import (
	"healthcare-capt-america/entities/models"
	"time"
)

type ChatResponse struct {
	ID          uint64    `json:"id"`
	SenderID    uint64    `json:"sender_id"`
	SenderName  string    `json:"sender_name"`
	SenderPhoto string    `json:"sender_photo"`
	Message     string    `json:"message"`
	SendTime    time.Time `json:"send_time"`
}

func NewChatResponse(chat *models.Chat) *ChatResponse {
	return &ChatResponse{
		ID:          chat.ID,
		SenderID:    chat.SenderID,
		SenderName:  chat.Sender.Name,
		SenderPhoto: chat.Sender.Photo,
		Message:     chat.Message,
		SendTime:    chat.CreatedAt,
	}
}

func NewChatResponses(chats []*models.Chat) []*ChatResponse {
	var responses []*ChatResponse
	for _, chat := range chats {
		responses = append(responses, NewChatResponse(chat))
	}
	return responses
}
