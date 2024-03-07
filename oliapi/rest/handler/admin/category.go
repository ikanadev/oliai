package admin

import (
	"net/http"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func postCategory(categoryRepo repository.CategoryRepository) echo.HandlerFunc {
	type requestData struct {
		BotID uuid.UUID `json:"botId" validate:"required,uuid4"`
		Name  string    `json:"name"  validate:"required,min=3,max=255"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		err := categoryRepo.SaveCategory(data.BotID, data.Name)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func getCategories(categoryRepo repository.CategoryRepository) echo.HandlerFunc {
	type requestData struct {
		BotID uuid.UUID `query:"botId" validate:"required,uuid4"`
	}

	return func(c echo.Context) error {
		var data requestData

		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		categories, err := categoryRepo.GetCategories(data.BotID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, categories)
	}
}

func updateCategory(categoryRepo repository.CategoryRepository) echo.HandlerFunc {
	type requestData struct {
		CategoryID uuid.UUID `param:"id"     validate:"required,uuid4"`
		Name       string    `json:"name"    validate:"required,min=3,max=255"`
		Archive    *bool     `json:"archive" validate:"required"`
		Delete     *bool     `json:"delete"  validate:"required"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		err := categoryRepo.UpdateCategory(repository.UpdateCategoryData{
			ID:      data.CategoryID,
			Name:    data.Name,
			Archive: *data.Archive,
			Delete:  *data.Delete,
		})
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
