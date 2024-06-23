package models

import "time"

// UserMembership represents the UserMembership model, joined User and Membership
type UserMembership struct {
	ID           int
	UserID       string
	MembershipID string
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
