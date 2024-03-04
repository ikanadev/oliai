package admin

import (
	"net/http"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/labstack/echo/v4"
)

func postCompany(companyRepo repository.CompanyRepository) echo.HandlerFunc {
	type requestData struct {
		Name    string `json:"name"    validate:"required,min=3,max=255"`
		LogoURL string `json:"logoUrl" validate:"required"`
	}

	return func(c echo.Context) error {
		var data requestData

		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		err := companyRepo.SaveCompany(data.Name, data.LogoURL)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, nil)
	}
}
