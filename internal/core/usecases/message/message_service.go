package message

import (
	"context"
	"errors"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/ports"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
	"github.com/hiago-balbino/hex-architecture-template/pkg/identifier"
)

type messageService struct {
	uuidGenerator identifier.UUIDGenerator
	repository    ports.MessageRepository
}

func NewMessageService(uuidGenerator identifier.UUIDGenerator, repository ports.MessageRepository) messageService {
	return messageService{
		uuidGenerator: uuidGenerator,
		repository:    repository,
	}
}

func (m messageService) Set(ctx context.Context, content string) (domain.Message, error) {
	message := domain.NewMessage(m.uuidGenerator.New(), content)
	err := m.repository.Set(ctx, message)
	if err != nil {
		return domain.Message{}, errors.Join(apperrors.InvalidInput, err)
	}
	return message, nil
}

func (m messageService) Get(ctx context.Context, id string) (domain.Message, error) {
	message, err := m.repository.Get(ctx, id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Message{}, err
		}
		return domain.Message{}, errors.Join(apperrors.InternalServerError, err)
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

func (m messageService) Delete(ctx context.Context, id string) error {
	err := m.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return err
		}
		return errors.Join(apperrors.InternalServerError, err)
	}
	return nil
}
