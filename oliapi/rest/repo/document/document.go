package document

import (
	"oliapi/db"
	"oliapi/domain"
	"oliapi/domain/repository"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewDocumentRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

type Repo struct {
	db *sqlx.DB
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
func (r Repo) SaveDocument(categoryID uuid.UUID, content string) error {
	now := time.Now()
	sql := `
		insert into documents (id, category_id, content, created_at, updated_at)
		values ($1, $2, $3, $4, $5);
	`

	_, err := r.db.Exec(sql, uuid.New(), categoryID, content, now, now)

	return err
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
