package models

import "time"

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

type PartialMembership struct {
	ID              int
	LibraryID       int
	Name            string
	DurationDays    int
	ActiveLoanLimit int
	FinePerDay      int

	Library *PartialLibrary
}
