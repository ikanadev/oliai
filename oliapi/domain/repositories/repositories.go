package repositories

import (
	"oliapi/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUser(id uuid.UUID) (domain.User, error)
}
