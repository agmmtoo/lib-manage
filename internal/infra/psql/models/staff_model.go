package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
)

// Staff represents Staff model, joined User and Library
type Staff struct {
	ID        int
	UserID    int
	LibraryID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// joined fields from User model
	UserUsername *string

	// joined fields from Library model
	LibraryName *string
}

func (s *Staff) ToCoreModel() *models.Staff {
	var u *models.PartialUser
	if s.UserUsername != nil {
		u = &models.PartialUser{
			ID:       s.UserID,
			Username: *s.UserUsername,
		}
	}

	var l *models.PartialLibrary
	if s.LibraryName != nil {
		l = &models.PartialLibrary{
			ID:   s.LibraryID,
			Name: *s.LibraryName,
		}
	}

	return &models.Staff{
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
