package public

import (
	"oliapi/domain/repository"

	"github.com/labstack/echo/v4"
)

func SetUpPublicRoutes(app *echo.Echo, userRepo repository.UserRepository, jwtKey []byte) {
	app.POST("/signup", signUp(userRepo))
	app.POST("/signin", signIn(userRepo, jwtKey))
}
