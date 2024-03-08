package repository

import (
	"oliapi/domain"

	"github.com/google/uuid"
)

type UpdateDocumentData struct {
	ID      uuid.UUID
	Content string
	Archive bool
	Delete  bool
}

type DocumentRepository interface {
	SaveDocument(categoryID uuid.UUID, content string) (uuid.UUID, error)
	GetDocuments(categoryID uuid.UUID) ([]domain.DocumentWithTimeData, error)
	GetDocument(documentID uuid.UUID) (domain.DocumentWithTimeData, error)
	UpdateDocument(data UpdateDocumentData) error
	GetBot(documentID uuid.UUID) (domain.BotWithTimeData, error)
}
