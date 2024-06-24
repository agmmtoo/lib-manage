package models

import "time"

type Category struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at,omitempty"`
}

type SubCategory struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	CategoryID int        `json:"category_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeleteAt   *time.Time `json:"deleted_at,omitempty"`

	Category *PartialCategory `json:"category"`
}

type PartialCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PartialSubCategory struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`

	Category *PartialCategory `json:"category"`
}
