package mocks

import (
	"context"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MessageServicerMock struct {
	mock.Mock
}

func (m *MessageServicerMock) Set(ctx context.Context, content string) (domain.Message, error) {
	args := m.Called(ctx, content)
	return args.Get(0).(domain.Message), args.Error(1)
}

func (m *MessageServicerMock) Get(ctx context.Context, id string) (domain.Message, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Message), args.Error(1)
}

func (m *MessageServicerMock) GetAll(ctx context.Context) ([]domain.Message, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *MessageServicerMock) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
