package models

import "time"

// Core Membership model
type Membership struct {
	ID              int        `json:"id"`
	LibraryID       int        `json:"library_id"`
	Name            string     `json:"name"`
	DurationDays    int        `json:"duration_days"`
	ActiveLoanLimit int        `json:"active_loan_limit"`
	FinePerDay      int        `json:"fine_per_day"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`

	Library *PartialLibrary `json:"library"`
}

type PartialMembership struct {
	ID              int    `json:"id"`
	LibraryID       int    `json:"library_id"`
	Name            string `json:"name"`
	DurationDays    int    `json:"duration_days"`
	ActiveLoanLimit int    `json:"active_loan_limit"`
	FinePerDay      int    `json:"fine_per_day"`
}
