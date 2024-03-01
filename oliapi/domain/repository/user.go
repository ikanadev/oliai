package repository

import (
	"oliapi/domain"

	"github.com/google/uuid"
)

type SaveUserData struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type UserRepository interface {
	GetUser(ID uuid.UUID) (domain.User, error)
	SaveUser(data SaveUserData) error
	VerifyUser(email, password string) (domain.User, error)
}
