package models

import "time"

// Core Library model
type Library struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type PartialLibrary struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
