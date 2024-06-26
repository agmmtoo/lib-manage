package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Core Membership model
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

	Library *PartialLibrary
}

func (m *Membership) ToAPIModel() *models.Membership {
	var l *models.PartialLibrary
	if m.Library != nil {
		l = &models.PartialLibrary{
			ID:   m.Library.ID,
			Name: m.Library.Name,
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

type PartialMembership struct {
	ID              int
	LibraryID       int
	Name            string
	DurationDays    int
	ActiveLoanLimit int
	FinePerDay      int

	Library *PartialLibrary
}
