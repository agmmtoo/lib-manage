package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Core Staff model
type Staff struct {
	ID        int
	UserID    int
	LibraryID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	User    *PartialUser
	Library *PartialLibrary
}

func (s Staff) ToAPIModel() models.Staff {
	var u *models.PartialUser
	if s.User != nil {
		u = &models.PartialUser{
			ID:       s.User.ID,
			Username: s.User.Username,
		}
	}

	var l *models.PartialLibrary
	if s.Library != nil {
		l = &models.PartialLibrary{
			ID:   s.Library.ID,
			Name: s.Library.Name,
		}
	}

	return models.Staff{
		ID:        s.ID,
		UserID:    s.UserID,
		LibraryID: s.LibraryID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		DeletedAt: s.DeletedAt,
		User:      u,
		Library:   l,
	}
}

type PartialStaff struct {
	ID        int
	UserID    int
	LibraryID int

	User    *PartialUser
	Library *PartialLibrary
}
