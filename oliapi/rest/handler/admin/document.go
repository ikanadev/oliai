package admin

import (
	"net/http"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func postDocument(documentRepository repository.DocumentRepository) echo.HandlerFunc {
	type requestData struct {
		CategoryID uuid.UUID `json:"categoryId" validate:"required,uuid4"`
		Content    string    `json:"content"    validate:"required"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		err := documentRepository.SaveDocument(data.CategoryID, data.Content)
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

func updateDocument(documentRepository repository.DocumentRepository) echo.HandlerFunc {
	type requestData struct {
		DocumentID uuid.UUID `param:"id"     validate:"required,uuid4"`
		Content    string    `json:"content" validate:"required"`
		Archive    *bool     `json:"archive" validate:"required"`
		Delete     *bool     `json:"delete"  validate:"required"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		err := documentRepository.UpdateDocument(repository.UpdateDocumentData{
			ID:      data.DocumentID,
			Content: data.Content,
			Archive: *data.Archive,
			Delete:  *data.Delete,
		})
		if err != nil {
			return err
		}

		return nil
	}
}
