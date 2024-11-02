package models

import "time"

type Category struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at,omitempty"`
}

type PartialCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
