package user

import (
	"context"
	"oliapi/domain"
	"oliapi/domain/repository"

	"github.com/google/uuid"
)

func NewUserRepo() Repo {
	return Repo{}
}

type Repo struct{}

// VerifyUser implements repository.UserRepository.
func (u Repo) VerifyUser(ctx context.Context, email string, password string) (domain.User, error) {
	var respUser domain.User

	return respUser, nil
}

// SaveUser implements repositories.UserRepository.
func (u Repo) SaveUser(ctx context.Context, data repository.SaveUserData) error {
	return nil
}

func (u Repo) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var appUser domain.User

	return appUser, nil
}
