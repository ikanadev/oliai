package rest

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type userTokenData struct {
	Id   string
	Role []string
}

func createToken(id string, jwtKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString(jwtKey)
}

func authMiddleware(jwtKey string) echo.MiddlewareFunc {
	return echojwt.JWT(jwtKey)
}
