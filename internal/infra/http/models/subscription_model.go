package models

import "time"

// API Subscription model
type Subscription struct {
	ID           int        `json:"id"`
	UserID       int        `json:"user_id"`
	MembershipID int        `json:"membership_id"`
	ExpiryDate   time.Time  `json:"expiry_date"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`

	User       *PartialUser       `json:"user"`
	Membership *PartialMembership `json:"membership"`
}

type PartialSubscription struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	MembershipID int       `json:"membership_id"`
	ExpiryDate   time.Time `json:"expiry_date"`

	User       *PartialUser       `json:"user"`
	Membership *PartialMembership `json:"membership"`
}
