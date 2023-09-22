package memory

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
)

var errNotFoundMessageID = errors.New("message id not found")

type messageStorage struct {
	data map[string][]byte
}

func NewMessageStorage() messageStorage {
	return messageStorage{
		data: make(map[string][]byte),
	}
}

func (m messageStorage) Save(ctx context.Context, message domain.Message) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	m.data[message.ID] = messageJSON
	return nil
}

func (m messageStorage) GetByID(ctx context.Context, id string) (domain.Message, error) {
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

func (m messageStorage) GetAll(ctx context.Context) ([]domain.Message, error) {
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

func (m messageStorage) DeleteByID(ctx context.Context, id string) error {
	if _, ok := m.data[id]; !ok {
		return errors.Join(apperrors.NotFound, errNotFoundMessageID)
	}
	delete(m.data, id)
	return nil
}
