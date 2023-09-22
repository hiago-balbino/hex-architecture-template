package ports

import (
	"context"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
)

type MessageUseCase interface {
	Save(ctx context.Context, content string) (domain.Message, error)
	GetByID(ctx context.Context, id string) (domain.Message, error)
	GetAll(ctx context.Context) ([]domain.Message, error)
	DeleteByID(ctx context.Context, id string) error
}
