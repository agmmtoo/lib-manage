package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
)

// Subscription represents the Subscription model, joined User and Membership
type Subscription struct {
	ID           int
	UserID       int
	MembershipID int
	ExpiryDate   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time

	// joined fields from User model
	UserUsername *string

	// joined fields from Membership model
	MembershipName            *string
	MembershipLibraryID       *int
	MembershipDurationDays    *int
	MembershipActiveLoanLimit *int
	MembershipFinePerDay      *int
	MembershipLibraryName     *string
}

func (s *Subscription) ToCoreModel() *models.Subscription {
	var u *models.PartialUser
	if s.UserUsername != nil {
		u = &models.PartialUser{
			ID:       s.UserID,
			Username: *s.UserUsername,
		}
	}

	var m *models.PartialMembership
	if s.MembershipName != nil {
		l := &models.PartialLibrary{
			ID:   *s.MembershipLibraryID,
			Name: *s.MembershipLibraryName,
		}

		m = &models.PartialMembership{
			ID:              s.MembershipID,
			LibraryID:       *s.MembershipLibraryID,
			Name:            *s.MembershipName,
			DurationDays:    *s.MembershipDurationDays,
			ActiveLoanLimit: *s.MembershipActiveLoanLimit,
			FinePerDay:      *s.MembershipFinePerDay,
			Library:         l,
		}
	}

	return &models.Subscription{
		ID:           s.ID,
		UserID:       s.UserID,
		MembershipID: s.MembershipID,
		ExpiryDate:   s.ExpiryDate,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
		DeletedAt:    s.DeletedAt,
		User:         u,
		Membership:   m,
	}
}
