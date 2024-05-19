package libraryapp

import (
	"time"
)

// Loan represents a loan of a book to a user.
type Loan struct {
	ID           int
	BookID       int
	UserID       int
	LibraryID    int
	StaffID      int
	LoanDate     time.Time
	DueDate      time.Time
	ReturnedDate *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
