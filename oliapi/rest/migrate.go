package rest

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
)

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
