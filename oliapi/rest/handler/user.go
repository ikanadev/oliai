package handler

import (
	"net/http"
	"oliapi/domain"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SetUpUserRoutes(app *echo.Group, userRepo repository.UserRepository, jwtKey []byte) {
	app.GET("/profile", getProfile(userRepo, jwtKey))
}

func getProfile(userRepo repository.UserRepository, jwtKey []byte) echo.HandlerFunc {
	type responseData struct {
		User  domain.User `json:"user"`
		Token string      `json:"token"`
	}

	return func(echoCtx echo.Context) error {
		token, ok := echoCtx.Get("id").(uuid.UUID)
		if !ok {
			return echo.ErrUnauthorized
		}

		user, err := userRepo.GetUser(token)
		if err != nil {
			return err
		}

		tokenStr, err := utils.CreateToken(user.ID, jwtKey)
		if err != nil {
			return err
		}

		return echoCtx.JSON(http.StatusOK, responseData{User: user, Token: tokenStr})
	}
}
