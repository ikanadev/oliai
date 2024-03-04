package utils

import "errors"

// unexpected errors.

var ErrCanNotParseToken = errors.New("can not parse token. Claims to MapClaims")

// expected errors.

const (
	ErrEmailAlreadyRegistered = "email_already_registered"
	ErrEmailIncorrect         = "email_incorrect"
	ErrEmailNotRegistered     = "email_not_registered"
	ErrPasswordIncorrect      = "password_incorrect"
)
