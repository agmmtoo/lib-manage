package models

import "time"

// API Book model
type Book struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	SubCategory *PartialSubCategory `json:"sub_category"`
}

type PartialBook struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	SubCategoryID *int   `json:"sub_category_id,omitempty"`

	SubCategory *PartialSubCategory `json:"sub_category,omitempty"`
}
