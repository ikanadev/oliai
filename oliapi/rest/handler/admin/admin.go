package admin

import (
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func SetUpAdminRoutes(app *echo.Group, companyRepo repository.CompanyRepository, db *sqlx.DB) {
	adminApp := app.Group("/admin")
	adminApp.Use(utils.AdminMiddleware(db))
	adminApp.POST("/company", postCompany(companyRepo))
	adminApp.GET("/company", getCompanies(companyRepo))
	adminApp.PUT("/company/:id", updateCompany(companyRepo))
}
