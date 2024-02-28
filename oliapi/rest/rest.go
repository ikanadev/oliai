package rest

import (
	"context"
	"net/http"
	"oliapi/domain/repository"
	"oliapi/ent"
	"oliapi/rest/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	// for postgres connection.
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
	ent      *ent.Client
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func (s Server) Start() {
	user.SetUpUserRoutes(s.app)

	panicIfError(s.app.Start(":" + s.config.Port))
}

func NewRestServer() Server {
	var server Server
	server.app = echo.New()
	server.app.Validator = &echoValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
	server.config = GetConfig()
	ent, err := ent.Open("postgres", server.config.DBConn)
	panicIfError(err)

	ctx := context.Background()
	err = ent.Schema.Create(ctx)
	panicIfError(err)
	err = populateStaticData(ctx, ent)
	panicIfError(err)

	server.ent = ent
	server.userRepo = user.NewUserRepo(server.ent)

	return server
}
