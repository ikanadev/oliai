package repository

import (
	"context"
	"oliapi/domain"

	"github.com/google/uuid"
)

type SaveVectorData struct {
	BotID      uuid.UUID
	DocumentID uuid.UUID
	Vector     []float32
}

type VectorRepository interface {
	CreateCollection(ctx context.Context, botID uuid.UUID, provider domain.EmbeddingProvider) error
	SaveVector(ctx context.Context, data SaveVectorData) error
}
