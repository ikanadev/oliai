package repository

import (
	"context"
	"oliapi/domain"
)

type EmbeddingRepository interface {
	EmbedContent(ctx context.Context, model domain.EmbeddingModel, content string) ([]float32, error)
}
