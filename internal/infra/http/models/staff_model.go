package models

import "time"

type Staff struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	LibraryID int        `json:"library_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	User    *PartialUser    `json:"user"`
	Library *PartialLibrary `json:"library"`
}

type PartialStaff struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	LibraryID int `json:"library_id"`

	User    *PartialUser    `json:"user"`
	Library *PartialLibrary `json:"library"`
}
