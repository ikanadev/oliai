package repository

import (
	"oliapi/domain"

	"github.com/google/uuid"
)

type UpdateCategoryData struct {
	ID      uuid.UUID
	Name    string
	Archive bool
	Delete  bool
}

type CategoryRepository interface {
	SaveCategory(botID uuid.UUID, name string) error
	GetCategories(botID uuid.UUID) ([]domain.CategoryWithTimeData, error)
	UpdateCategory(data UpdateCategoryData) error
}
