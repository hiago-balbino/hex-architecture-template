package ports

import (
	"context"

	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
)

type MessageServicer interface {
	Set(ctx context.Context, content string) (domain.Message, error)
	Get(ctx context.Context, id string) (domain.Message, error)
	GetAll(ctx context.Context) ([]domain.Message, error)
}
