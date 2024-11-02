package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
)

// Setting represents base Setting model
type Setting struct {
	ID        int
	LibraryID int
	Key       string
	Value     string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// joined fields from Library model
	LibraryName *string
}

func (s Setting) ToCoreModel() models.Setting {
	var l *models.PartialLibrary
	if s.LibraryName != nil {
		l = &models.PartialLibrary{
			ID:   s.LibraryID,
			Name: *s.LibraryName,
		}
	}

	return models.Setting{
		ID:        s.ID,
		LibraryID: s.LibraryID,
		Key:       s.Key,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		DeletedAt: s.DeletedAt,
		Library:   l,
	}
}
