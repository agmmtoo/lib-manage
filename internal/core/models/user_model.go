package models

import "time"

// Core User model
type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PartialUser struct {
	ID       int
	Username string
}
