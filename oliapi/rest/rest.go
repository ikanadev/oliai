package rest

import (
	"context"
	"oliapi/domain/repositories"
	"oliapi/ent"
	"oliapi/rest/user"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type RestServer struct {
	app      *echo.Echo
	config   Config
	userRepo repositories.UserRepository
	ent      *ent.Client
}

func (s RestServer) Start() error {
	user.SetUpUserRoutes(s.app)

	return s.app.Start(":" + s.config.Port)
}

func NewRestServer() (RestServer, error) {
	server := RestServer{}
	server.app = echo.New()
	server.config = GetConfig()
	ent, err := ent.Open("postgres", server.config.DbConn)
	if err != nil {
		return server, err
	}
	ctx := context.Background()
	err = ent.Schema.Create(ctx)
	if err != nil {
		return server, err
	}
	err = populateStaticData(ent, ctx)
	if err != nil {
		return server, err
	}
	server.ent = ent
	server.userRepo = user.NewUserRepo(server.ent)
	return server, nil
}
