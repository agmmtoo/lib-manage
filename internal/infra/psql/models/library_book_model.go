package models

import "time"

// LibraryBook represents LibraryBook model, joined Library and Book
type LibraryBook struct {
	ID        int
	LibraryID int
	BookID    int
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// joined fields from Library model
	LibraryName *string

	// joined fields from Book model
	BookTitle                   *string
	BookAuthor                  *string
	BookSubCategoryID           *int
	BookSubCategoryName         *string
	BookSubCategoryCategoryID   *int
	BookSubCategoryCategoryName *string
}
