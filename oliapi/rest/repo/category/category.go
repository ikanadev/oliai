package category

import (
	"oliapi/db"
	"oliapi/domain"
	"oliapi/domain/repository"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewCategoryRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

type Repo struct {
	db *sqlx.DB
}

// GetCategories implements repository.CategoryRepository.
func (r Repo) GetCategories(botID uuid.UUID) ([]domain.CategoryWithTimeData, error) {
	var dbCategories []db.Category

	err := r.db.Select(&dbCategories, "select * from categories where bot_id=$1", botID)
	if err != nil {
		return nil, err
	}

	categories := make([]domain.CategoryWithTimeData, len(dbCategories))

	for i := range dbCategories {
		dbCat := dbCategories[i]
		categories[i] = domain.CategoryWithTimeData{
			Category: domain.Category{
				ID:   dbCat.ID,
				Name: dbCat.Name,
			},
			TimeData: domain.TimeData{
				CreatedAt:  dbCat.CreatedAt,
				UpdatedAt:  dbCat.UpdatedAt,
				DeletedAt:  dbCat.DeletedAt,
				ArchivedAt: dbCat.ArchivedAt,
			},
		}
	}

	return categories, nil
}

// SaveCategory implements repository.CategoryRepository.
func (r Repo) SaveCategory(botID uuid.UUID, name string) error {
	now := time.Now()
	sql := `
		insert into categories (id, bot_id, name, created_at, updated_at)
		values ($1, $2, $3, $4, $5);
`
	_, err := r.db.Exec(sql, uuid.New(), botID, name, now, now)

	return err
}

// UpdateCategory implements repository.CategoryRepository.
func (r Repo) UpdateCategory(data repository.UpdateCategoryData) error {
	now := time.Now()
	deletedAt := &now
	archivedAt := &now
	sql := `
		update categories set name=$1, archived_at=$2, deleted_at=$3, updated_at=$4 where id=$5;
`

	if !data.Delete {
		deletedAt = nil
	}

	if !data.Archive {
		archivedAt = nil
	}

	_, err := r.db.Exec(
		sql,
		data.Name,
		archivedAt,
		deletedAt,
		now,
		data.ID,
	)

	return err
}
