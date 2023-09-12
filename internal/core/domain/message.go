package domain

import "github.com/google/uuid"

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func NewMessage(content string) Message {
	return Message{
		ID:      uuid.New().String(),
		Content: content,
	}
}
