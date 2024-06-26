package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
)

// Library model
type Library struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (l *Library) ToCoreModel() *models.Library {
	return &models.Library{
		ID:        l.ID,
		Name:      l.Name,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
		DeletedAt: l.DeletedAt,
	}
}
