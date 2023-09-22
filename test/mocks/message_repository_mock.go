package mocks

import (
	"context"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MessageRepositoryMock struct {
	mock.Mock
}

func (m *MessageRepositoryMock) Save(ctx context.Context, message domain.Message) error {
	args := m.Called(ctx, message)
	return args.Error(0)
}

func (m *MessageRepositoryMock) GetByID(ctx context.Context, id string) (domain.Message, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Message), args.Error(1)
}

func (m *MessageRepositoryMock) GetAll(ctx context.Context) ([]domain.Message, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *MessageRepositoryMock) DeleteByID(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
