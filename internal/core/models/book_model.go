package models

import "time"

// Core Book model
type Book struct {
	ID            int
	Title         string
	Author        string
	SubCategoryID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time

	SubCategory *PartialSubCategory
}

type PartialBook struct {
	ID            int
	Title         string
	Author        string
	SubCategoryID int

	SubCategory *PartialSubCategory
}

// Core Book model
type LibraryBook struct {
	ID        int
	Code      string
	LibraryID int
	BookID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Book    *PartialBook
	Library *PartialLibrary
}

type PartialLibraryBook struct {
	ID        int
	Code      string
	LibraryID int
	BookID    int
	Book      *PartialBook
	Library   *PartialLibrary
}
