package models

import "time"

// Core UserMembership model
type UserMembership struct {
	ID           int
	UserID       int
	MembershipID int
	ExpiryDate   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time

	User       *PartialUser
	Membership *PartialMembership
}
type PartialUserMembership struct {
	ID           int
	UserID       int
	MembershipID int
	ExpiryDate   time.Time

	User       *PartialUser
	Membership *PartialMembership
}
