package user

import (
	"net/http"
	"oliapi/domain"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/labstack/echo/v4"
)

func SetUpUserRoutes(app *echo.Echo, userRepo repository.UserRepository) {
	app.POST("/signup", signUp(userRepo))
	app.POST("/signin", signIn(userRepo))
}

func signUp(userRepo repository.UserRepository) echo.HandlerFunc {
	type requestData struct {
		Email     string `json:"email"     validate:"required,email"`
		FirstName string `json:"firstName" validate:"required,min=3,max=255"`
		LastName  string `json:"lastName"  validate:"required,min=3,max=255"`
		Password  string `json:"password"  validate:"required,min=8,max=255"`
	}

	return func(echoCtx echo.Context) error {
		var data requestData
		if err := echoCtx.Bind(&data); err != nil {
			return utils.NewRestErr(err)
		}

		if err := echoCtx.Validate(data); err != nil {
			return err
		}

		err := userRepo.SaveUser(echoCtx.Request().Context(), repository.SaveUserData{
			Email:     data.Email,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Password:  data.Password,
		})
		if err != nil {
			return err
		}

		return echoCtx.JSON(http.StatusCreated, nil)
	}
}

func signIn(userRepo repository.UserRepository) echo.HandlerFunc {
	type requestData struct {
		Email    string `json:"email"    validate:"required,email"`
		Password string `json:"password" validate:"required,max=255"`
	}
	type responseData struct {
		User  domain.User `json:"user"`
		Token string      `json:"token"`
	}

	return func(echoCtx echo.Context) error {
		var data requestData
		if err := echoCtx.Bind(&data); err != nil {
			return utils.NewRestErr(err)
		}

		if err := echoCtx.Validate(data); err != nil {
			return err
		}

		user, err := userRepo.VerifyUser(echoCtx.Request().Context(), data.Email, data.Password)
		if err != nil {
			return err
		}

		return echoCtx.JSON(http.StatusCreated, user)
	}
}
