package public

import (
	"net/http"
	"oliapi/domain"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/labstack/echo/v4"
)

func signUp(userRepo repository.UserRepository) echo.HandlerFunc {
	type requestData struct {
		Email     string `json:"email"     validate:"required,email"`
		FirstName string `json:"firstName" validate:"required,min=3,max=255"`
		LastName  string `json:"lastName"  validate:"required,min=3,max=255"`
		Password  string `json:"password"  validate:"required,min=8,max=255"`
	}

	return func(c echo.Context) error {
		var data requestData
		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		err := userRepo.SaveUser(repository.SaveUserData{
			Email:     data.Email,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Password:  data.Password,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

func signIn(userRepo repository.UserRepository, jwtKey []byte) echo.HandlerFunc {
	type requestData struct {
		Email    string `json:"email"    validate:"required,email"`
		Password string `json:"password" validate:"required,max=255"`
	}

	type responseData struct {
		User  domain.User `json:"user"`
		Token string      `json:"token"`
	}

	return func(c echo.Context) error {
		var data requestData

		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		user, err := userRepo.VerifyUser(data.Email, data.Password)
		if err != nil {
			return err
		}

		token, err := utils.CreateToken(user.ID, jwtKey)
		if err != nil {
			return err
		}

		resp := responseData{
			User:  user,
			Token: token,
		}

		return c.JSON(http.StatusCreated, resp)
	}
}
