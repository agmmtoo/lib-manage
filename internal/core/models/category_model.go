package models

import "time"

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}

type SubCategory struct {
	ID         int
	Name       string
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   *time.Time

	Category *PartialCategory
}

type PartialCategory struct {
	ID   int
	Name string
}

type PartialSubCategory struct {
	ID         int
	CategoryID int
	Name       string

	Category *PartialCategory
}
