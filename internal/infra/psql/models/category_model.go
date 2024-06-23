package models

import "time"

// Category model
type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// SubCategory model, joined Category
type SubCategory struct {
	ID         int
	Name       string
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time

	// joined fields from Category model
	CategoryName *string
}
