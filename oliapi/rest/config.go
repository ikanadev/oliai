package rest

import "os"

type Config struct {
	Port          string
	JWTKey        []byte
	DBConn        string
	MigrationsURL string
	QDrantURL     string
}

//nolint:gochecknoglobals
var _config = Config{
	Port:          os.Getenv("PORT"),
	JWTKey:        []byte(os.Getenv("JWT_SIGN_KEY")),
	DBConn:        os.Getenv("DATABASE"),
	MigrationsURL: os.Getenv("MIGRATIONS_URL"),
	QDrantURL:     os.Getenv("QDRANT_URL"),
}

func GetConfig() Config {
	return _config
}
