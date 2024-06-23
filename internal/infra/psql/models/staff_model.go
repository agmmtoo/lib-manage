package models

import "time"

// Staff represents Staff model, joined User and Library
type Staff struct {
	ID        int
	UserID    string
	LibraryID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// joined fields from User model
	UserUsername *string

	// joined fields from Library model
	LibraryName *string
}
