package models

import "time"

type Setting struct {
	ID        int        `json:"id"`
	LibraryID int        `json:"library_id"`
	Key       string     `json:"key"`
	Value     string     `json:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Library *PartialLibrary `json:"library"`
}
