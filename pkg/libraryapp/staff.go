package libraryapp

import "time"

// Staff represents a staff member of the library.
type Staff struct {
	ID     int
	UserID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
