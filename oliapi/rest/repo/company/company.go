package company

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewCompanyRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

type Repo struct {
	db *sqlx.DB
}

// SaveCompany implements repository.AdminRepository.
func (r Repo) SaveCompany(name string, logoURL string) error {
	now := time.Now()
	sql := `
insert into companies (id, name, logo_url, created_at, updated_at)
values ($1, $2, $3, $4, $5);
`
	_, err := r.db.Exec(sql, uuid.New(), name, logoURL, now, now)

	return err
}
