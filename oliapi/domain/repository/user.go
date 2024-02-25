package repository

import (
	"context"
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
	GetUser(ctx context.Context, ID uuid.UUID) (domain.User, error)
	SaveUser(ctx context.Context, data SaveUserData) error
	VerifyUser(ctx context.Context, email, password string) (domain.User, error)
}
