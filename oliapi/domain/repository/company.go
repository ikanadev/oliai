package repository

type CompanyRepository interface {
	SaveCompany(name string, logoURL string) error
}
