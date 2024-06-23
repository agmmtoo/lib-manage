package models

import "time"

// Core Library model
type Library struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PartialLibrary struct {
	ID   int
	Name string
}
