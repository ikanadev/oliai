package bot

import (
	"oliapi/db"
	"oliapi/domain"
	"oliapi/domain/repository"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewBotRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

type Repo struct {
	db *sqlx.DB
}

// GetBots implements repository.BotRepository.
func (r Repo) GetBots(companyID uuid.UUID) ([]domain.BotWithTimeData, error) {
	var dbBots []db.Bot

	err := r.db.Select(&dbBots, "select * from bots where company_id = $1;", companyID)
	if err != nil {
		return nil, err
	}

	bots := make([]domain.BotWithTimeData, len(dbBots))

	for i := range dbBots {
		dbBot := dbBots[i]
		bots[i] = domain.BotWithTimeData{
			Bot: domain.Bot{
				ID:              dbBot.ID,
				Name:            dbBot.Name,
				EmbeddingModel:  domain.EmbeddingModel(dbBot.EmbeddingModel),
				GreetingMessage: dbBot.GreetingMessage,
				CustomPrompt:    dbBot.CustomPrompt,
			},
			TimeData: domain.TimeData{
				CreatedAt:  dbBot.CreatedAt,
				UpdatedAt:  dbBot.UpdatedAt,
				DeletedAt:  dbBot.DeletedAt,
				ArchivedAt: dbBot.ArchivedAt,
			},
		}
	}

	return bots, nil
}

// SaveBot implements repository.BotRepository.
func (r Repo) SaveBot(name string, companyID uuid.UUID, model domain.EmbeddingModel) (uuid.UUID, error) {
	const (
		defaultGreetingMessage = "Hola, ¿cómo puedo ayudarte?"
		defaultCustomPrompt    = "Eres un asistente virtual."
	)

	uuid := uuid.New()
	now := time.Now()
	sql := `
		INSERT INTO bots (id, company_id, name, embedding_model, greeting_message, custom_prompt, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	_, err := r.db.Exec(sql, uuid, companyID, name, model, defaultGreetingMessage, defaultCustomPrompt, now, now)

	return uuid, err
}

// UpdateBot implements repository.BotRepository.
func (r Repo) UpdateBot(data repository.UpdateBotData) error {
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
		UPDATE bots
		SET name = $1, greeting_message = $2, custom_prompt = $3, archived_at = $4, deleted_at = $5, updated_at = $6
		WHERE id = $7;
	`

	_, err := r.db.Exec(
		sql,
		data.Name,
		data.GreetingMessage,
		data.CustomPrompt,
		archivedAt,
		deletedAt,
		now,
		data.ID,
	)

	return err
}
