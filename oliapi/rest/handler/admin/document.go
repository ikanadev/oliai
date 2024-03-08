package admin

import (
	"context"
	"net/http"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func postDocument(
	documentRepository repository.DocumentRepository,
	vectorRepo repository.VectorRepository,
	embeddingRepo repository.EmbeddingRepository,
) echo.HandlerFunc {
	type requestData struct {
		CategoryID uuid.UUID `json:"categoryId" validate:"required,uuid4"`
		Content    string    `json:"content"    validate:"required"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		documentID, err := documentRepository.SaveDocument(data.CategoryID, data.Content)
		if err != nil {
			return err
		}

		err = embeddDocument(c.Request().Context(), documentID, data.Content, documentRepository, embeddingRepo, vectorRepo)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func getDocuments(documentRepository repository.DocumentRepository) echo.HandlerFunc {
	type requestData struct {
		CategoryID uuid.UUID `query:"categoryId" validate:"required,uuid4"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		documents, err := documentRepository.GetDocuments(data.CategoryID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, documents)
	}
}

func updateDocument(
	documentRepository repository.DocumentRepository,
	embeddingRepo repository.EmbeddingRepository,
	vectorRepo repository.VectorRepository,
) echo.HandlerFunc {
	type requestData struct {
		DocumentID uuid.UUID `param:"id"     validate:"required,uuid4"`
		Content    string    `json:"content" validate:"required"`
		Archive    *bool     `json:"archive" validate:"required"`
		Delete     *bool     `json:"delete"  validate:"required"`
	}

	return func(c echo.Context) error {
		var (
			data         requestData
			shouldEmbedd = false
		)

		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		doc, err := documentRepository.GetDocument(data.DocumentID)
		if err != nil {
			return err
		}

		if doc.Content != data.Content {
			shouldEmbedd = true
		}

		err = documentRepository.UpdateDocument(repository.UpdateDocumentData{
			ID:      data.DocumentID,
			Content: data.Content,
			Archive: *data.Archive,
			Delete:  *data.Delete,
		})
		if err != nil {
			return err
		}

		if shouldEmbedd {
			err := embeddDocument(c.Request().Context(), doc.ID, data.Content, documentRepository, embeddingRepo, vectorRepo)
			if err != nil {
				return err
			}
		}

		return c.NoContent(http.StatusOK)
	}
}

func embeddDocument(
	ctx context.Context,
	documentID uuid.UUID,
	content string,
	documentRepository repository.DocumentRepository,
	embeddingRepo repository.EmbeddingRepository,
	vectorRepo repository.VectorRepository,
) error {
	bot, err := documentRepository.GetBot(documentID)
	if err != nil {
		return err
	}

	vector, err := embeddingRepo.EmbedContent(ctx, bot.EmbeddingModel, content)
	if err != nil {
		return err
	}

	err = vectorRepo.SaveVector(ctx, repository.SaveVectorData{
		BotID:      bot.ID,
		DocumentID: documentID,
		Vector:     vector,
	})
	if err != nil {
		return err
	}

	return nil
}
