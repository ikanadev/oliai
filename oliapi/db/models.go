package db

import (
	"time"

	"github.com/google/uuid"
)

type TimeFields struct {
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	ArchivedAt *time.Time `db:"archived_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}

type User struct {
	ID        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Password  string    `db:"password"`
	TimeFields
}

type Role struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type UserRole struct {
	UserID uuid.UUID `db:"user_id"`
	RoleID uuid.UUID `db:"role_id"`
}

type Company struct {
	ID      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	LogoURL string    `db:"logo_url"`
	TimeFields
}

type UserCompany struct {
	UserID    uuid.UUID `db:"user_id"`
	CompanyID uuid.UUID `db:"company_id"`
}

type Bot struct {
	ID              uuid.UUID `db:"id"`
	CompanyID       uuid.UUID `db:"company_id"`
	Name            string    `db:"name"`
	GreetingMessage string    `db:"greeting_message"`
	CustomPrompt    string    `db:"custom_prompt"`
	EmbeddingModel  string    `db:"embedding_model"`
	TimeFields
}

type Category struct {
	ID    uuid.UUID `db:"id"`
	BotID uuid.UUID `db:"bot_id"`
	Name  string    `db:"name"`
	TimeFields
}

type Document struct {
	ID            uuid.UUID  `db:"id"`
	CategoryID    uuid.UUID  `db:"category_id"`
	Content       string     `db:"content"`
	EmbeddingDate *time.Time `db:"embedding_date"`
	TimeFields
}
