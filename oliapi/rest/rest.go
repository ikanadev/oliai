package rest

import (
	"errors"
	"oliapi/domain/repository"
	"oliapi/rest/handler/admin"
	"oliapi/rest/handler/common"
	"oliapi/rest/handler/public"
	"oliapi/rest/repo/bot"
	"oliapi/rest/repo/company"
	"oliapi/rest/repo/user"
	"oliapi/rest/utils"

	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	// for postgres connection.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
)

func NewRestServer() Server {
	var server Server
	server.config = GetConfig()
	server.db = sqlx.MustConnect("postgres", server.config.DBConn)
	server.app = echo.New()
	server.app.Debug = true
	server.app.Validator = &echoValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
	server.protectedApp = server.app.Group("/api")
	server.protectedApp.Use(echojwt.JWT(server.config.JWTKey), utils.UserIDMiddleware)

	// repositories
	server.userRepo = user.NewUserRepo(server.db)
	server.companyRepo = company.NewCompanyRepo(server.db)
	server.botRepo = bot.NewBotRepo(server.db)

	return server
}

type Server struct {
	app          *echo.Echo
	protectedApp *echo.Group
	config       Config
	userRepo     repository.UserRepository
	companyRepo  repository.CompanyRepository
	botRepo      repository.BotRepository
	db           *sqlx.DB
}

func (s Server) Migrate() {
	migrator, err := migrate.New(s.config.MigrationsURL, s.config.DBConn)
	panicIfError(err)

	err = migrator.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			panicIfError(err)
		}
	}
}

func (s Server) Start() {
	public.SetUpPublicRoutes(s.app, s.userRepo, s.config.JWTKey)
	common.SetUpCommonRoutes(s.protectedApp, s.userRepo, s.config.JWTKey)
	admin.SetUpAdminRoutes(s.protectedApp, s.companyRepo, s.botRepo, s.db)
	panicIfError(s.app.Start(":" + s.config.Port))
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
