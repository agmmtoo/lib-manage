package models

import "time"

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
