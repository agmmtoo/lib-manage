package libraryapp

import "time"

// LibraryStaff represents a staff member's association with a library.
type LibraryStaff struct {
	LibraryID int
	StaffID   int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
