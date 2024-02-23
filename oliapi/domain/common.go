package domain

import "time"

type TimeData struct {
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ArchivedAt *time.Time `json:"archivedAt"`
}
