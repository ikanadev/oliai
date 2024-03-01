package utils

import "errors"

type ClientError struct {
	Status int
	Err    error
}

func (c ClientError) Error() string {
	return c.Err.Error()
}

func NewClientErr(status int, err error) ClientError {
	return ClientError{
		Status: status,
		Err:    err,
	}
}

type RestError struct {
	Err error
}

func (r RestError) Error() string {
	return r.Err.Error()
}

func NewRestErr(err error) RestError {
	return RestError{
		Err: err,
	}
}

// unexpected errors.

var ErrCanNotParseToken = errors.New("can not parse token. Claims to MapClaims")

// expected errors.

const (
	ErrEmailAlreadyRegistered = "email_already_registered"
	ErrEmailIncorrect         = "email_incorrect"
	ErrEmailNotRegistered     = "email_not_registered"
	ErrPasswordIncorrect      = "password_incorrect"
)
