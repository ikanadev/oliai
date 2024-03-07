package rest

import (
	"oliapi/domain/repository"
	"oliapi/rest/handler/admin"
	"oliapi/rest/handler/common"
	"oliapi/rest/handler/public"
	"oliapi/rest/repo/bot"
	"oliapi/rest/repo/category"
	"oliapi/rest/repo/company"
	"oliapi/rest/repo/document"
	"oliapi/rest/repo/user"
	"oliapi/rest/repo/vector"
	"oliapi/rest/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// for postgres connection.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
)

func NewRestServer() Server {
	var server Server
	// config
	server.config = GetConfig()
	// db config
	server.db = sqlx.MustConnect("postgres", server.config.DBConn)
	// echo app
	server.app = echo.New()
	server.app.Debug = true
	server.app.Validator = &echoValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
	server.protectedApp = server.app.Group("/api")
	server.protectedApp.Use(echojwt.JWT(server.config.JWTKey), utils.UserIDMiddleware)
	// Qdrant
	conn, err := grpc.Dial(server.config.QDrantURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	panicIfError(err)

	server.grpc = conn
	// repositories
	server.userRepo = user.NewUserRepo(server.db)
	server.companyRepo = company.NewCompanyRepo(server.db)
	server.botRepo = bot.NewBotRepo(server.db)
	server.categoryRepo = category.NewCategoryRepo(server.db)
	server.documentRepo = document.NewDocumentRepo(server.db)
	server.vectorRepo = vector.NewVectorRepo(server.qdrant)

	return server
}

type Server struct {
	app          *echo.Echo
	protectedApp *echo.Group
	config       Config
	userRepo     repository.UserRepository
	companyRepo  repository.CompanyRepository
	botRepo      repository.BotRepository
	categoryRepo repository.CategoryRepository
	documentRepo repository.DocumentRepository
	vectorRepo   repository.VectorRepository
	db           *sqlx.DB
	grpc         *grpc.ClientConn
}

func (s Server) Start() {
	public.SetUpPublicRoutes(s.app, s.userRepo, s.config.JWTKey)
	common.SetUpCommonRoutes(s.protectedApp, s.userRepo, s.config.JWTKey)
	admin.SetUpAdminRoutes(s.protectedApp, s.companyRepo, s.botRepo, s.categoryRepo, s.documentRepo, s.db)
	panicIfError(s.app.Start(":" + s.config.Port))
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
