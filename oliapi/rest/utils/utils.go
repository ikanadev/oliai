package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID string, jwtKey string) (string, error) {
	const validHours = 72

	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", ErrCanNotParseToken
	}

	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * validHours).Unix()

	tokenStr, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", NewRestErr(err)
	}

	return tokenStr, nil
}

func AuthMiddleware(jwtKey string) echo.MiddlewareFunc {
	return echojwt.JWT(jwtKey)
}
