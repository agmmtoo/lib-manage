package models

import "time"

// Core Loan model
type Loan struct {
	ID               int
	LibraryBookID    int
	UserMembershipID int
	StaffID          int
	LoanDate         time.Time
	DueDate          time.Time
	ReturnDate       time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time

	LibraryBook    *PartialLibraryBook
	UserMembership *PartialUserMembership
}
