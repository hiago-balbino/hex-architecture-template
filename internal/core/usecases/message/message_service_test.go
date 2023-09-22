package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
	"github.com/hiago-balbino/hex-architecture-template/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSave_ShouldReturnErrorWhenRepositoryFails(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()
	content := "message content"
	unexpectedError := errors.New("unexpected error")

	identifierMock := new(mocks.UUIDGeneratorMock)
	repositoryMock := new(mocks.MessageRepositoryMock)
	identifierMock.On("New").Return(messageID)
	repositoryMock.On("Save", ctx, domain.NewMessage(messageID, content)).Return(unexpectedError)

	service := NewMessageService(identifierMock, repositoryMock)
	actualMessage, err := service.Save(ctx, content)

	assert.ErrorIs(t, err, apperrors.InvalidInput)
	assert.Empty(t, actualMessage)
}

func TestSave_ShouldSaveMessageWithSuccess(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()
	content := "message content"

	identifierMock := new(mocks.UUIDGeneratorMock)
	repositoryMock := new(mocks.MessageRepositoryMock)
	identifierMock.On("New").Return(messageID)
	repositoryMock.On("Save", ctx, domain.NewMessage(messageID, content)).Return(nil)

	service := NewMessageService(identifierMock, repositoryMock)
	actualMessage, err := service.Save(ctx, content)

	assert.NoError(t, err)
	assert.Equal(t, messageID, actualMessage.ID)
	assert.Equal(t, content, actualMessage.Content)
}

func TestGetByID_ShouldReturnErrorWhenRepositoryFails(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()
	unexpectedError := errors.New("unexpected error")

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("GetByID", ctx, messageID).Return(domain.Message{}, unexpectedError)

	service := NewMessageService(nil, repositoryMock)
	actualMessage, err := service.GetByID(ctx, messageID)

	assert.ErrorIs(t, err, apperrors.InternalServerError)
	assert.Empty(t, actualMessage)
}

func TestGetByID_ShouldReturnErrorWhenMessageNotFound(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("GetByID", ctx, messageID).Return(domain.Message{}, apperrors.NotFound)

	service := NewMessageService(nil, repositoryMock)
	actualMessage, err := service.GetByID(ctx, messageID)

	assert.ErrorIs(t, err, apperrors.NotFound)
	assert.Empty(t, actualMessage)
}

func TestGetByID_ShouldReturnMessageWithSuccess(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()
	expectedMessage := domain.NewMessage(messageID, "message content")

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("GetByID", ctx, messageID).Return(expectedMessage, nil)

	service := NewMessageService(nil, repositoryMock)
	actualMessage, err := service.GetByID(ctx, messageID)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, actualMessage)
}

func TestGetAll_ShouldReturnErrorWhenRepositoryFails(t *testing.T) {
	ctx := context.Background()
	unexpectedError := errors.New("unexpected error")

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("GetAll", ctx).Return([]domain.Message{}, unexpectedError)

	service := NewMessageService(nil, repositoryMock)
	actualMessages, err := service.GetAll(ctx)

	assert.ErrorIs(t, err, apperrors.InternalServerError)
	assert.Empty(t, actualMessages)
}

func TestGetAll_ShouldReturnAllMessagesWithSuccess(t *testing.T) {
	ctx := context.Background()
	firstMessage := domain.NewMessage("id1", "message content 1")
	secondMessage := domain.NewMessage("id2", "message content 2")
	expectedMessages := []domain.Message{firstMessage, secondMessage}

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("GetAll", ctx).Return(expectedMessages, nil)

	service := NewMessageService(nil, repositoryMock)
	actualMessages, err := service.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, actualMessages)
}

func TestDeleteByID_ShouldReturnErrorWhenRepositoryFails(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()
	unexpectedError := errors.New("unexpected error")

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("DeleteByID", ctx, messageID).Return(unexpectedError)

	service := NewMessageService(nil, repositoryMock)
	err := service.DeleteByID(ctx, messageID)

	assert.ErrorIs(t, err, apperrors.InternalServerError)
}

func TestDeleteByID_ShouldReturnErrorWhenMessageNotFound(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("DeleteByID", ctx, messageID).Return(apperrors.NotFound)

	service := NewMessageService(nil, repositoryMock)
	err := service.DeleteByID(ctx, messageID)

	assert.ErrorIs(t, err, apperrors.NotFound)
}

func TestDeleteByID_ShouldDeleteMessageWithSuccess(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()

	repositoryMock := new(mocks.MessageRepositoryMock)
	repositoryMock.On("DeleteByID", ctx, messageID).Return(nil)

	service := NewMessageService(nil, repositoryMock)
	err := service.DeleteByID(ctx, messageID)

	assert.NoError(t, err)
}
