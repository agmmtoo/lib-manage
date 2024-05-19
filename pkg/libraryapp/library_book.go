package libraryapp

import "time"

// LibraryBook represents a book's association with a library.
type LibraryBook struct {
	LibraryID int
	BookID    int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
