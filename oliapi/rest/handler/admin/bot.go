package admin

import (
	"net/http"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func postBot(botRepo repository.BotRepository) echo.HandlerFunc {
	type requestData struct {
		Name      string `json:"name" validate:"required,min=3,max=255"`
		CompanyID string `param:"id"  validate:"required,uuid4"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		companyID, err := uuid.Parse(data.CompanyID)
		if err != nil {
			return err
		}

		err = botRepo.SaveBot(data.Name, companyID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func updateBot(botRepo repository.BotRepository) echo.HandlerFunc {
	type requestData struct {
		BotID           string `param:"bot_id"         validate:"required,uuid4"`
		Name            string `json:"name"            validate:"required,min=3,max=255"`
		GreetingMessage string `json:"greetingMessage" validate:"required"`
		CustomPrompt    string `json:"customPrompt"    validate:"required"`
		Archive         *bool  `json:"archive"         validate:"required"`
		Delete          *bool  `json:"delete"          validate:"required"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		botID, err := uuid.Parse(data.BotID)
		if err != nil {
			return err
		}

		err = botRepo.UpdateBot(repository.UpdateBotData{
			ID:              botID,
			Name:            data.Name,
			GreetingMessage: data.GreetingMessage,
			CustomPrompt:    data.CustomPrompt,
			Archive:         *data.Archive,
			Delete:          *data.Delete,
		})
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func getBots(botRepo repository.BotRepository) echo.HandlerFunc {
	type requestData struct {
		CompanyID string `param:"id" validate:"required,uuid4"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		companyID, err := uuid.Parse(data.CompanyID)
		if err != nil {
			return err
		}

		bots, err := botRepo.GetBots(companyID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, bots)
	}
}
