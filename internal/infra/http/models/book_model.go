package models

import "time"

// Core Book model
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

// Core Book model
type LibraryBook struct {
	ID        int        `json:"id"`
	Code      string     `json:"code"`
	LibraryID int        `json:"library_id"`
	BookID    int        `json:"book_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Book    *PartialBook    `json:"book"`
	Library *PartialLibrary `json:"library"`
}

type PartialLibraryBook struct {
	ID        int             `json:"id"`
	Code      string          `json:"code"`
	LibraryID int             `json:"library_id"`
	BookID    int             `json:"book_id"`
	Book      *PartialBook    `json:"book"`
	Library   *PartialLibrary `json:"library"`
}
