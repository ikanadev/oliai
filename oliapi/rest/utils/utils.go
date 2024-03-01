package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID uuid.UUID, jwtKey []byte) (string, error) {
	const validHours = 72

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID.String(),
		"iss": "oliai",
		"exp": time.Now().Add(time.Hour * validHours).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", NewRestErr(err)
	}

	return tokenStr, nil
}

func AuthMiddleware(jwtKey string) echo.MiddlewareFunc {
	return echojwt.JWT(jwtKey)
}
