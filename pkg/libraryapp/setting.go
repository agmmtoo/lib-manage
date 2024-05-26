package libraryapp

import "time"

// Setting represents a setting of a library.
type Setting struct {
	ID        int
	LibraryID int
	Key       string
	Value     string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
