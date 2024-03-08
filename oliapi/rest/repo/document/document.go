package document

import (
	"database/sql"
	"errors"
	"oliapi/db"
	"oliapi/domain"
	"oliapi/domain/repository"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func NewDocumentRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

type Repo struct {
	db *sqlx.DB
}

// GetDocument implements repository.DocumentRepository.
func (r Repo) GetDocument(documentID uuid.UUID) (domain.DocumentWithTimeData, error) {
	var (
		dbDocument db.Document
		document   domain.DocumentWithTimeData
	)

	sqlQuery := "select * from documents where id = $1;"

	err := r.db.Get(&dbDocument, sqlQuery, documentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return document, echo.ErrNotFound
		}

		return document, err
	}

	document = domain.DocumentWithTimeData{
		Document: domain.Document{
			ID:      dbDocument.ID,
			Content: dbDocument.Content,
		},
		TimeData: domain.TimeData{
			CreatedAt:  dbDocument.CreatedAt,
			UpdatedAt:  dbDocument.UpdatedAt,
			DeletedAt:  dbDocument.DeletedAt,
			ArchivedAt: dbDocument.ArchivedAt,
		},
	}

	return document, nil
}

// GetBot implements repository.DocumentRepository.
func (r Repo) GetBot(documentID uuid.UUID) (domain.BotWithTimeData, error) {
	var dbBot db.Bot

	var bot domain.BotWithTimeData

	sql := `
		select * from bots where id in (
			select bot_id from categories where id in (
				select category_id from documents where id = $1
			)
		);
	`

	err := r.db.Get(&dbBot, sql, documentID)
	if err != nil {
		return bot, err
	}

	bot = domain.BotWithTimeData{
		Bot: domain.Bot{
			ID:              dbBot.ID,
			Name:            dbBot.Name,
			GreetingMessage: dbBot.GreetingMessage,
			CustomPrompt:    dbBot.CustomPrompt,
			EmbeddingModel:  domain.EmbeddingModel(dbBot.EmbeddingModel),
		},
		TimeData: domain.TimeData{
			CreatedAt:  dbBot.CreatedAt,
			UpdatedAt:  dbBot.UpdatedAt,
			DeletedAt:  dbBot.DeletedAt,
			ArchivedAt: dbBot.ArchivedAt,
		},
	}

	return bot, nil
}

// GetDocuments implements repository.DocumentRepository.
func (r Repo) GetDocuments(categoryID uuid.UUID) ([]domain.DocumentWithTimeData, error) {
	var dbDocuments []db.Document

	err := r.db.Select(&dbDocuments, "select * from documents where category_id=$1;", categoryID)
	if err != nil {
		return nil, err
	}

	documents := make([]domain.DocumentWithTimeData, len(dbDocuments))

	for i := range dbDocuments {
		dbDoc := dbDocuments[i]
		documents[i] = domain.DocumentWithTimeData{
			Document: domain.Document{
				ID:      dbDoc.ID,
				Content: dbDoc.Content,
			},
			TimeData: domain.TimeData{
				CreatedAt:  dbDoc.CreatedAt,
				UpdatedAt:  dbDoc.UpdatedAt,
				DeletedAt:  dbDoc.DeletedAt,
				ArchivedAt: dbDoc.ArchivedAt,
			},
		}
	}

	return documents, nil
}

// SaveDocument implements repository.DocumentRepository.
func (r Repo) SaveDocument(categoryID uuid.UUID, content string) (uuid.UUID, error) {
	now := time.Now()
	id := uuid.New()
	sql := `
		insert into documents (id, category_id, content, created_at, updated_at)
		values ($1, $2, $3, $4, $5);
	`

	_, err := r.db.Exec(sql, id, categoryID, content, now, now)

	return id, err
}

// UpdateDocument implements repository.DocumentRepository.
func (r Repo) UpdateDocument(data repository.UpdateDocumentData) error {
	now := time.Now()
	deletedAt := &now
	archivedAt := &now
	sql := `
		update documents set content=$1, archived_at=$2, deleted_at=$3, updated_at=$4 where id=$5;
`

	if !data.Delete {
		deletedAt = nil
	}

	if !data.Archive {
		archivedAt = nil
	}

	_, err := r.db.Exec(
		sql,
		data.Content,
		archivedAt,
		deletedAt,
		now,
		data.ID,
	)

	return err
}
