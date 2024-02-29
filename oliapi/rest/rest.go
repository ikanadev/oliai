package rest

import (
	"errors"
	"net/http"
	"oliapi/domain/repository"
	"oliapi/rest/user"

	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	// for postgres connection.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
)

type echoValidator struct {
	validator *validator.Validate
}

func (e echoValidator) Validate(i any) error {
	if err := e.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

type Server struct {
	app      *echo.Echo
	config   Config
	userRepo repository.UserRepository
	db       *sqlx.DB
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
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
	user.SetUpUserRoutes(s.app, s.userRepo)
	panicIfError(s.app.Start(":" + s.config.Port))
}

func NewRestServer() Server {
	var server Server
	server.app = echo.New()
	server.app.Debug = true
	server.app.Validator = &echoValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
	server.config = GetConfig()
	server.db = sqlx.MustConnect("postgres", server.config.DBConn)
	server.userRepo = user.NewUserRepo()

	return server
}
