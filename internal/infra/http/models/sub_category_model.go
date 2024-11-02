package models

import "time"

type SubCategory struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	CategoryID int        `json:"category_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeleteAt   *time.Time `json:"deleted_at,omitempty"`

	Category *PartialCategory `json:"category"`
}

type PartialSubCategory struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`

	Category *PartialCategory `json:"category"`
}
