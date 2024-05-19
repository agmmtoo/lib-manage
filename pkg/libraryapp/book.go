package libraryapp

import "time"

// Book represents a book in the library.
type Book struct {
	ID     int
	Title  string
	Author string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
