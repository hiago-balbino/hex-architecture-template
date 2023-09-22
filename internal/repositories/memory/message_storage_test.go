package memory

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
	"github.com/stretchr/testify/assert"
)

func TestSave_ShouldSaveMessageWithSuccess(t *testing.T) {
	message := domain.NewMessage("id", "message content")

	repo := NewMessageStorage()
	err := repo.Save(context.Background(), message)

	assert.NoError(t, err)
}

func TestGetByID_ShouldReturnErrorWhenMessageNotFound(t *testing.T) {
	messageID := uuid.NewString()

	repo := NewMessageStorage()
	actualMessage, err := repo.GetByID(context.Background(), messageID)

	assert.ErrorIs(t, err, apperrors.NotFound)
	assert.Empty(t, actualMessage)
}

func TestGetByID_ShouldReturnErrorWhenInvalidMessageContent(t *testing.T) {
	messageID := uuid.NewString()
	invalidMessageContent := []byte("{")

	repo := messageStorage{data: map[string][]byte{messageID: invalidMessageContent}}
	message, err := repo.GetByID(context.Background(), messageID)

	assert.Error(t, err)
	assert.Empty(t, message)
}

func TestGetByID_ShouldGetMessageWithSuccess(t *testing.T) {
	ctx := context.Background()
	expectedMessage := domain.NewMessage("id", "message content")

	repo := NewMessageStorage()
	err := repo.Save(ctx, expectedMessage)
	assert.NoError(t, err)

	actualMessage, err := repo.GetByID(ctx, expectedMessage.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, actualMessage)
}

func TestGetAll_ShouldReturnErrorWhenInvalidMessageContent(t *testing.T) {
	ctx := context.Background()
	firstMessage := domain.Message{ID: uuid.NewString(), Content: "message content 1"}
	secondMessage := domain.Message{ID: uuid.NewString(), Content: "{"}

	repo := messageStorage{data: map[string][]byte{
		firstMessage.ID:  []byte(firstMessage.Content),
		secondMessage.ID: []byte(secondMessage.Content),
	}}
	messages, err := repo.GetAll(ctx)

	assert.Error(t, err)
	assert.Empty(t, messages)
}

func TestGetAll_ShouldReturnAllMessagesWithSuccess(t *testing.T) {
	ctx := context.Background()
	firstMessage := domain.NewMessage("id1", "message content 1")
	secondMessage := domain.NewMessage("id2", "message content 2")
	expectedMessages := []domain.Message{firstMessage, secondMessage}

	repo := NewMessageStorage()
	err := repo.Save(ctx, firstMessage)
	assert.NoError(t, err)
	err = repo.Save(ctx, secondMessage)
	assert.NoError(t, err)

	actualMessages, err := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, actualMessages)
}

func TestDeleteByID_ShouldReturnErrorWhenMessageNotFound(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()

	repo := NewMessageStorage()
	err := repo.DeleteByID(ctx, messageID)

	assert.ErrorIs(t, err, apperrors.NotFound)
}

func TestDeleteByID_ShouldDeleteMessageWithSuccess(t *testing.T) {
	ctx := context.Background()
	message := domain.NewMessage("id", "message content")

	repo := NewMessageStorage()
	err := repo.Save(ctx, message)
	assert.NoError(t, err)

	err = repo.DeleteByID(ctx, message.ID)
	assert.NoError(t, err)
}
