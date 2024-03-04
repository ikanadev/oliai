package domain

import "github.com/google/uuid"

type Company struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	LogoURL string    `json:"logoUrl"`
}

type CompanyWithTimeData struct {
	Company
	TimeData
}

type CompanyWithBots struct {
	Company
	Bots []BotWithCategories `json:"bots"`
}
