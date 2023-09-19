package domain

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func NewMessage(messageID string, content string) Message {
	return Message{
		ID:      messageID,
		Content: content,
	}
}
