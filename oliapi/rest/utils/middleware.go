package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Must be used after the echo-jwt middleware since it assumes that the token is already valid.
func UserIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		token, ok := echoCtx.Get("user").(*jwt.Token)
		if !ok {
			return echo.ErrUnauthorized
		}

		userIDStr, err := token.Claims.GetSubject()
		if err != nil {
			return err
		}

		id, err := uuid.Parse(userIDStr)
		if err != nil {
			return err
		}

		echoCtx.Set("id", id)

		return next(echoCtx)
	}
}
