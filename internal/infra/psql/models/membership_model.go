package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
)

// Membership represents the Membership model, joined Library
type Membership struct {
	ID              int
	LibraryID       int
	Name            string
	DurationDays    int
	ActiveLoanLimit int
	FinePerDay      int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time

	// joined fields from Library model
	LibraryName *string
}

func (m *Membership) ToCoreModel() *models.Membership {
	var l *models.PartialLibrary
	if m.LibraryName != nil {
		l = &models.PartialLibrary{
			ID:   m.LibraryID,
			Name: *m.LibraryName,
		}
	}

	return &models.Membership{
		ID:              m.ID,
		LibraryID:       m.LibraryID,
		Name:            m.Name,
		DurationDays:    m.DurationDays,
		ActiveLoanLimit: m.ActiveLoanLimit,
		FinePerDay:      m.FinePerDay,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		DeletedAt:       m.DeletedAt,
		Library:         l,
	}
}
