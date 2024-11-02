package models

import "time"

// API LibraryBook model
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
