package rest

import "os"

type Config struct {
	Port   string
	JWTKey []byte
	DbConn string
}

var _config = Config{
	Port:   os.Getenv("PORT"),
	JWTKey: []byte(os.Getenv("JWT_SIGN_KEY")),
	DbConn: os.Getenv("DATABASE"),
}

func GetConfig() Config {
	return _config
}
