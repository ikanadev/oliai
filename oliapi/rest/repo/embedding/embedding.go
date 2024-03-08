package embedding

import (
	"context"
	"oliapi/domain"

	"github.com/labstack/echo/v4"
)

func NewEmbeddingRepo() Repo {
	return Repo{}
}

type Repo struct{}

// EmbedContent implements repository.EmbeddingRepository.
func (r Repo) EmbedContent(ctx context.Context, model domain.EmbeddingModel, content string) ([]float32, error) {
	if true {
		return make([]float32, 1536), nil
	}

	switch model {
	case domain.EmbeddingModelOpenAI3Small:
		return embedWithOpenAIText3Small(ctx, content)
	case domain.EmbeddingModelOpenAI3Large:
		return embedWithOpenAIText3Large(ctx, content)
	}

	return nil, echo.ErrInternalServerError
}
