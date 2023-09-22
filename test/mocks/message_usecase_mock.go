package mocks

import (
	"context"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MessageUseCaseMock struct {
	mock.Mock
}

func (m *MessageUseCaseMock) Save(ctx context.Context, content string) (domain.Message, error) {
	args := m.Called(ctx, content)
	return args.Get(0).(domain.Message), args.Error(1)
}

func (m *MessageUseCaseMock) GetByID(ctx context.Context, id string) (domain.Message, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Message), args.Error(1)
}

func (m *MessageUseCaseMock) GetAll(ctx context.Context) ([]domain.Message, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *MessageUseCaseMock) DeleteByID(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
