package messagerepo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
	"github.com/stretchr/testify/assert"
)

func TestSet_ShouldSetMessageWithSuccess(t *testing.T) {
	message := domain.NewMessage("message content")

	repo := NewMessageRepository()
	err := repo.Set(context.Background(), message)

	assert.NoError(t, err)
}

func TestGet_ShouldReturnErrorWhenMessageNotFound(t *testing.T) {
	messageID := uuid.NewString()

	repo := NewMessageRepository()
	actualMessage, err := repo.Get(context.Background(), messageID)

	assert.ErrorIs(t, err, apperrors.NotFound)
	assert.Empty(t, actualMessage)
}

func TestGet_ShouldReturnErrorWhenInvalidMessageContent(t *testing.T) {
	messageID := uuid.NewString()
	invalidMessageContent := []byte("{")

	repo := messageRepository{data: map[string][]byte{messageID: invalidMessageContent}}
	message, err := repo.Get(context.Background(), messageID)

	assert.Error(t, err)
	assert.Empty(t, message)
}

func TestGet_ShouldGetMessageWithSuccess(t *testing.T) {
	ctx := context.Background()
	expectedMessage := domain.NewMessage("message content")

	repo := NewMessageRepository()
	err := repo.Set(ctx, expectedMessage)
	assert.NoError(t, err)

	actualMessage, err := repo.Get(ctx, expectedMessage.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, actualMessage)
}

func TestGetAll_ShouldReturnErrorWhenInvalidMessageContent(t *testing.T) {
	ctx := context.Background()
	firstMessage := domain.Message{ID: uuid.NewString(), Content: "message 1"}
	secondMessage := domain.Message{ID: uuid.NewString(), Content: "{"}

	repo := messageRepository{data: map[string][]byte{
		firstMessage.ID:  []byte(firstMessage.Content),
		secondMessage.ID: []byte(secondMessage.Content),
	}}
	messages, err := repo.GetAll(ctx)

	assert.Error(t, err)
	assert.Empty(t, messages)
}

func TestGetAll_ShouldReturnAllMessagesWithSuccess(t *testing.T) {
	ctx := context.Background()
	firstMessage := domain.NewMessage("message 1")
	secondMessage := domain.NewMessage("message 2")
	expectedMessages := []domain.Message{firstMessage, secondMessage}

	repo := NewMessageRepository()
	err := repo.Set(ctx, firstMessage)
	assert.NoError(t, err)
	err = repo.Set(ctx, secondMessage)
	assert.NoError(t, err)

	actualMessages, err := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, actualMessages)
}

func TestDelete_ShouldReturnErrorWhenMessageNotFound(t *testing.T) {
	ctx := context.Background()
	messageID := uuid.NewString()

	repo := NewMessageRepository()
	err := repo.Delete(ctx, messageID)

	assert.ErrorIs(t, err, apperrors.NotFound)
}

func TestDelete_ShouldDeleteMessageWithSuccess(t *testing.T) {
	ctx := context.Background()
	message := domain.NewMessage("message content")

	repo := NewMessageRepository()
	err := repo.Set(ctx, message)
	assert.NoError(t, err)

	err = repo.Delete(ctx, message.ID)
	assert.NoError(t, err)
}
