package dto

import "github.com/hiago-balbino/hex-architecture-template/internal/core/domain"

type GetMessageResponse domain.Message

func BuildResponseGetMessage(message domain.Message) GetMessageResponse {
	return GetMessageResponse(message)
}

func BuildResponseGetMessages(messages []domain.Message) []GetMessageResponse {
	messagesDto := []GetMessageResponse{}
	for _, message := range messages {
		messagesDto = append(messagesDto, BuildResponseGetMessage(message))
	}
	return messagesDto
}
