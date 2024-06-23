package models

import "time"

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

type PartialUser struct {
	ID       int
	Username string
}
