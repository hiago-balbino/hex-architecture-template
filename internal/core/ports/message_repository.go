package ports

import (
	"context"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
)

type MessageRepository interface {
	Set(ctx context.Context, message domain.Message) error
	Get(ctx context.Context, id string) (domain.Message, error)
	GetAll(ctx context.Context) ([]domain.Message, error)
	Delete(ctx context.Context, id string) error
}
