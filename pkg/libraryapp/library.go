package libraryapp

import "time"

// Library represents a library.
type Library struct {
	ID   int
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
