package message

import (
	"context"
	"errors"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/ports"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
)

type messageService struct {
	repository ports.MessageRepository
}

func NewMessageService(repository ports.MessageRepository) messageService {
	return messageService{
		repository: repository,
	}
}

func (m messageService) Set(ctx context.Context, content string) (domain.Message, error) {
	message := domain.NewMessage(content)
	err := m.repository.Set(ctx, message)
	if err != nil {
		return domain.Message{}, errors.Join(apperrors.InvalidInput, err)
	}
	return message, nil
}

func (m messageService) Get(ctx context.Context, id string) (domain.Message, error) {
	message, err := m.repository.Get(ctx, id)
	if err != nil {
		return domain.Message{}, errors.Join(apperrors.NotFound, err)
	}
	return message, nil
}

func (m messageService) GetAll(ctx context.Context) ([]domain.Message, error) {
	messages, err := m.repository.GetAll(ctx)
	if err != nil {
		return nil, errors.Join(apperrors.InternalServerError, err)
	}
	return messages, nil
}
