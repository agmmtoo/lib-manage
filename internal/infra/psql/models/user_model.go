package models

import "time"

// User model
type User struct {
	ID        int
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
