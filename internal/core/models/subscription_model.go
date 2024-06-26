package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Core Subscription model
type Subscription struct {
	ID           int
	UserID       int
	MembershipID int
	ExpiryDate   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time

	User       *PartialUser
	Membership *PartialMembership
}

func (s *Subscription) ToAPIModel() *models.Subscription {
	var u *models.PartialUser
	if s.User != nil {
		u = &models.PartialUser{
			ID:       s.User.ID,
			Username: s.User.Username,
		}
	}

	var m *models.PartialMembership
	if s.Membership != nil {

		var l *models.PartialLibrary
		if s.Membership.Library != nil {
			l = &models.PartialLibrary{
				ID:   s.Membership.Library.ID,
				Name: s.Membership.Library.Name,
			}
		}

		m = &models.PartialMembership{
			ID:              s.Membership.ID,
			LibraryID:       s.Membership.LibraryID,
			Name:            s.Membership.Name,
			DurationDays:    s.Membership.DurationDays,
			ActiveLoanLimit: s.Membership.ActiveLoanLimit,
			FinePerDay:      s.Membership.FinePerDay,
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

type PartialSubscription struct {
	ID           int
	UserID       int
	MembershipID int
	ExpiryDate   time.Time

	User       *PartialUser
	Membership *PartialMembership
}
