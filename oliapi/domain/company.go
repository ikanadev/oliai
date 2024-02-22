package domain

type Company struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	LogoUrl string `json:"logoUrl"`
	Email   string `json:"email"`
}
