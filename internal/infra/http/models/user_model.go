package models

import "time"

// API User model
type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Password string `json:"password"`
}

type PartialUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
