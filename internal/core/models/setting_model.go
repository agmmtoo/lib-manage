package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Core Setting model
type Setting struct {
	ID        int
	LibraryID int
	Key       string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Library *PartialLibrary
}

func (s Setting) ToAPIModel() models.Setting {
	var l *models.PartialLibrary
	if s.Library != nil {
		pl := s.Library.ToAPIModel()
		l = &pl
	}
	return models.Setting{
		ID:        s.ID,
		LibraryID: s.LibraryID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		DeletedAt: s.DeletedAt,
		Library:   l,
	}
}
