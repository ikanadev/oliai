package repository

import (
	"oliapi/domain"

	"github.com/google/uuid"
)

type UpdateBotData struct {
	ID              uuid.UUID
	Name            string
	GreetingMessage string
	CustomPrompt    string
	Delete          bool
	Archive         bool
}

type BotRepository interface {
	SaveBot(name string, companyID uuid.UUID) error
	GetBots(companyID uuid.UUID) ([]domain.BotWithTimeData, error)
	UpdateBot(data UpdateBotData) error
}
