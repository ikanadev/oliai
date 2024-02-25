package utils

import "errors"

var ErrCanNotParseToken = errors.New("can not parse token. Claims to MapClaims")

type ClientError struct {
	Status  int
	Message string
}

func (c ClientError) Error() string {
	return c.Message
}

func NewClientErr(status int, message string) ClientError {
	return ClientError{
		Status:  status,
		Message: message,
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
