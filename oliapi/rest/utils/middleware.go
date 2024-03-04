package utils

import (
	"oliapi/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

// Must be used after the UserIDMiddleware.
func AdminMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			var roles []string

			userID, ok := echoCtx.Get("id").(uuid.UUID)
			if !ok {
				return echo.ErrUnauthorized
			}

			sql := `
select name from roles where id in (
	select role_id from users_roles where user_id = $1
);
`

			err := db.Select(&roles, sql, userID)
			if err != nil {
				return err
			}

			isAdmin := false

			for i := range roles {
				if domain.RoleAdmin == domain.RoleFromSting(roles[i]) {
					isAdmin = true

					break
				}
			}

			if !isAdmin {
				return echo.ErrUnauthorized
			}

			return next(echoCtx)
		}
	}
}
