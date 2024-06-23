package models

import "time"

// Loan represents the Loan model, joined LibraryBook, UserMembership, Staff
type Loan struct {
	ID               int
	LibraryBookID    int
	UserMembershipID int
	StaffID          int
	LoanDate         time.Time
	DueDate          time.Time
	ReturnDate       *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time

	// joined fields from Staff model
	StaffUserID       *int
	StaffUserUsername *string
	StaffLibraryID    *int
	StaffLibraryName  *string

	// joined fields from LibraryBook model
	LibraryBookBookID                      *int
	LibraryBookLibraryID                   *int
	LibraryBookBookCode                    *string
	LibraryBookLibraryName                 *string
	LibraryBookBookTitle                   *string
	LibraryBookBookAuthor                  *string
	LibraryBookBookSubCategoryID           *int
	LibraryBookBookSubCategoryName         *string
	LibraryBookBookSubCategoryCategoryID   *int
	LibraryBookBookSubCategoryCategoryName *string

	// joined fields from UserMembership model
	UserMembershipUserID                    *int
	UserMembershipExpiryDate                *time.Time
	UserMembershipUserUsername              *string
	UserMembershipMembershipName            *string
	UserMembershipMembershipLibraryID       *int
	UserMembershipMembershipDurationDays    *int
	UserMembershipMembershipActiveLoanLimit *int
	UserMembershipMembershipFinePerDay      *int
	UserMembershipMembershipLibraryName     *string
}
