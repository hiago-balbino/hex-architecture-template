package messagerepo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
)

var errNotFoundMessageID = errors.New("message id not found")

type messageRepository struct {
	data map[string][]byte
}

func NewMessageRepository() messageRepository {
	return messageRepository{
		data: make(map[string][]byte),
	}
}

func (m messageRepository) Set(ctx context.Context, message domain.Message) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	m.data[message.ID] = messageJSON
	return nil
}

func (m messageRepository) Get(ctx context.Context, id string) (domain.Message, error) {
	messageJSON, ok := m.data[id]
	if !ok {
		return domain.Message{}, errors.Join(apperrors.NotFound, errNotFoundMessageID)
	}

	var message domain.Message
	err := json.Unmarshal(messageJSON, &message)
	if err != nil {
		return domain.Message{}, err
	}

	return message, nil
}

func (m messageRepository) GetAll(ctx context.Context) ([]domain.Message, error) {
	var messages []domain.Message

	for _, messageJSON := range m.data {
		var message domain.Message
		err := json.Unmarshal(messageJSON, &message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (m messageRepository) Delete(ctx context.Context, id string) error {
	if _, ok := m.data[id]; !ok {
		return errors.Join(apperrors.NotFound, errNotFoundMessageID)
	}
	delete(m.data, id)
	return nil
}
