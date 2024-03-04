package company

import (
	"oliapi/db"
	"oliapi/domain"
	"oliapi/domain/repository"
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

// GetCompaniesWithTimeData implements repository.CompanyRepository.
func (r Repo) GetCompaniesWithTimeData() ([]domain.CompanyWithTimeData, error) {
	var dbCompanies []db.Company

	err := r.db.Select(&dbCompanies, "select * from companies;")
	if err != nil {
		return nil, err
	}

	companies := make([]domain.CompanyWithTimeData, len(dbCompanies))

	for i := range dbCompanies {
		dbComp := dbCompanies[i]
		companies[i] = domain.CompanyWithTimeData{
			Company: domain.Company{
				ID:      dbComp.ID,
				Name:    dbComp.Name,
				LogoURL: dbComp.LogoURL,
			},
			TimeData: domain.TimeData{
				CreatedAt:  dbComp.CreatedAt,
				UpdatedAt:  dbComp.UpdatedAt,
				DeletedAt:  dbComp.DeletedAt,
				ArchivedAt: dbComp.ArchivedAt,
			},
		}
	}

	return companies, nil
}

// UpdateCompany implements repository.CompanyRepository.
func (r Repo) UpdateCompany(data repository.UpdateCompanyData) error {
	now := time.Now()
	deletedAt := &now
	archivedAt := &now

	if !data.Delete {
		deletedAt = nil
	}

	if !data.Archive {
		archivedAt = nil
	}

	sql := `
update companies set name=$1, logo_url=$2, deleted_at=$3, archived_at=$4, updated_at=$5 where id=$6;
`

	_, err := r.db.Exec(
		sql,
		data.Name,
		data.LogoURL,
		deletedAt,
		archivedAt,
		now,
		data.ID,
	)
	if err != nil {
		return err
	}

	return nil
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
