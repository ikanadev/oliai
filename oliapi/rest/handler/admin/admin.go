package admin

import (
	"oliapi/domain/repository"
	"oliapi/rest/utils"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func SetUpAdminRoutes(
	app *echo.Group,
	companyRepo repository.CompanyRepository,
	botRepo repository.BotRepository,
	categoryRepo repository.CategoryRepository,
	documentRepo repository.DocumentRepository,
	vectorRepo repository.VectorRepository,
	embeddingRepo repository.EmbeddingRepository,
	db *sqlx.DB,
) {
	adminApp := app.Group("/admin")
	adminApp.Use(utils.AdminMiddleware(db))
	adminApp.POST("/companies", postCompany(companyRepo))
	adminApp.GET("/companies", getCompanies(companyRepo))
	adminApp.PUT("/companies/:id", updateCompany(companyRepo))
	adminApp.POST("/bots", postBot(botRepo, vectorRepo))
	adminApp.GET("/bots", getBots(botRepo))
	adminApp.PUT("/bots/:id", updateBot(botRepo))
	adminApp.POST("/categories", postCategory(categoryRepo))
	adminApp.GET("/categories", getCategories(categoryRepo))
	adminApp.PUT("/categories/:id", updateCategory(categoryRepo))
	adminApp.POST("/documents", postDocument(documentRepo, vectorRepo, embeddingRepo))
	adminApp.GET("/documents", getDocuments(documentRepo))
	adminApp.PUT("/documents/:id", updateDocument(documentRepo, embeddingRepo, vectorRepo))
}
