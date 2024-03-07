package admin

import (
	"net/http"
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/google/uuid"
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

func getCompanies(companyRepo repository.CompanyRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		companies, err := companyRepo.GetCompaniesWithTimeData()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, companies)
	}
}

func updateCompany(companyRepo repository.CompanyRepository) echo.HandlerFunc {
	type requestData struct {
		ID      uuid.UUID `param:"id"     validate:"required,uuid4"`
		Name    string    `json:"name"    validate:"required,min=3,max=255"`
		LogoURL string    `json:"logoUrl" validate:"required"`
		Archive *bool     `json:"archive" validate:"required"`
		Delete  *bool     `json:"delete"  validate:"required"`
	}

	return func(c echo.Context) error {
		var data requestData

		if err := utils.BindAndValidate(c, &data); err != nil {
			return err
		}

		err := companyRepo.UpdateCompany(repository.UpdateCompanyData{
			ID:      data.ID,
			Name:    data.Name,
			LogoURL: data.LogoURL,
			Archive: *data.Archive,
			Delete:  *data.Delete,
		})
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
