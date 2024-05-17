package libraryapp

import "time"

// User represents a user of the library.
type User struct {
	ID       int
	Username string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
