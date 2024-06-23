package models

import "time"

// Book represents base Book model
type Book struct {
	ID            int
	Title         string
	Author        string
	SubCategoryID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// joined fields from SubCategory model
	SubCategoryName         *string
	SubCategoryCategoryID   *int
	SubCategoryCategoryName *string
}
