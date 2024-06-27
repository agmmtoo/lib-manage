package models

import "time"

// API Loan model
type Loan struct {
	ID             int        `json:"id"`
	LibraryBookID  int        `json:"library_book_id"`
	SubscriptionID int        `json:"subscription_id"`
	StaffID        int        `json:"staff_id"`
	LoanDate       time.Time  `json:"loan_date"`
	DueDate        time.Time  `json:"due_date"`
	ReturnDate     *time.Time `json:"return_date,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`

	Staff        *PartialStaff        `json:"staff"`
	LibraryBook  *PartialLibraryBook  `json:"library_book"`
	Subscription *PartialSubscription `json:"subscription"`
}
