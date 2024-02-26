package domain

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"name"`
	LastName  string    `json:"lastName"`
	TimeData
}
