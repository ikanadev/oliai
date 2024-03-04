package repository

import (
	"oliapi/domain"

	"github.com/google/uuid"
)

type UpdateCompanyData struct {
	ID      uuid.UUID
	Name    string
	LogoURL string
	Archive bool
	Delete  bool
}

type CompanyRepository interface {
	SaveCompany(name string, logoURL string) error
	UpdateCompany(data UpdateCompanyData) error
	GetCompaniesWithTimeData() ([]domain.CompanyWithTimeData, error)
}
