package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
)

// Loan represents the Loan model, joined LibraryBook, Subscription, Staff
type Loan struct {
	ID             int
	LibraryBookID  int
	SubscriptionID int
	StaffID        int
	LoanDate       time.Time
	DueDate        time.Time
	ReturnDate     *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time

	// joined fields from Staff model
	StaffUserID       *int
	StaffUserUsername *string
	StaffLibraryID    *int
	StaffLibraryName  *string

	// joined fields from LibraryBook model
	LibraryBookLibraryID                   *int
	LibraryBookBookID                      *int
	LibraryBookCode                        *string
	LibraryBookLibraryName                 *string
	LibraryBookBookTitle                   *string
	LibraryBookBookAuthor                  *string
	LibraryBookBookSubCategoryID           *int
	LibraryBookBookSubCategoryName         *string
	LibraryBookBookSubCategoryCategoryID   *int
	LibraryBookBookSubCategoryCategoryName *string

	// joined fields from Subscription model
	SubscriptionUserID                    *int
	SubscriptionMembershipID              *int
	SubscriptionExpiryDate                *time.Time
	SubscriptionUserUsername              *string
	SubscriptionMembershipName            *string
	SubscriptionMembershipLibraryID       *int
	SubscriptionMembershipDurationDays    *int
	SubscriptionMembershipActiveLoanLimit *int
	SubscriptionMembershipFinePerDay      *int
	SubscriptionMembershipLibraryName     *string
}

func (l *Loan) ToCoreModel() *models.Loan {

	var staff *models.PartialStaff
	if l.StaffUserID != nil {

		var user *models.PartialUser
		if l.StaffUserUsername != nil {
			user = &models.PartialUser{
				ID:       *l.StaffUserID,
				Username: *l.StaffUserUsername,
			}
		}

		var library *models.PartialLibrary
		if l.StaffLibraryID != nil {
			library = &models.PartialLibrary{
				ID:   *l.StaffLibraryID,
				Name: *l.StaffLibraryName,
			}
		}

		staff = &models.PartialStaff{
			ID:        l.StaffID,
			UserID:    *l.StaffUserID,
			LibraryID: *l.StaffLibraryID,
			User:      user,
			Library:   library,
		}
	}

	var libraryBook *models.PartialLibraryBook
	if l.LibraryBookLibraryID != nil {

		var book *models.PartialBook
		if l.LibraryBookBookID != nil {

			var sc *models.PartialSubCategory
			if l.LibraryBookBookSubCategoryID != nil {

				var c *models.PartialCategory
				if l.LibraryBookBookSubCategoryCategoryID != nil {
					c = &models.PartialCategory{
						ID:   *l.LibraryBookBookSubCategoryCategoryID,
						Name: *l.LibraryBookBookSubCategoryCategoryName,
					}
				}
				sc = &models.PartialSubCategory{
					ID:         *l.LibraryBookBookSubCategoryID,
					Name:       *l.LibraryBookBookSubCategoryName,
					CategoryID: *l.LibraryBookBookSubCategoryCategoryID,
					Category:   c,
				}
			}

			book = &models.PartialBook{
				ID:            *l.LibraryBookBookID,
				Title:         *l.LibraryBookBookTitle,
				Author:        *l.LibraryBookBookAuthor,
				SubCategoryID: l.LibraryBookBookSubCategoryID,
				SubCategory:   sc,
			}
		}

		var lib *models.PartialLibrary
		if l.LibraryBookLibraryName != nil {
			lib = &models.PartialLibrary{
				ID:   *l.LibraryBookLibraryID,
				Name: *l.LibraryBookLibraryName,
			}
		}

		libraryBook = &models.PartialLibraryBook{
			ID:        l.LibraryBookID,
			Code:      *l.LibraryBookCode,
			LibraryID: *l.LibraryBookLibraryID,
			BookID:    *l.LibraryBookBookID,
			Book:      book,
			Library:   lib,
		}
	}

	var subscription *models.PartialSubscription
	if l.SubscriptionUserID != nil {
		var user *models.PartialUser
		if l.SubscriptionUserID != nil {
			user = &models.PartialUser{
				ID:       *l.SubscriptionUserID,
				Username: *l.SubscriptionUserUsername,
			}
		}

		var membership *models.PartialMembership
		if l.SubscriptionMembershipID != nil {

			var library *models.PartialLibrary
			if l.SubscriptionMembershipLibraryID != nil {
				library = &models.PartialLibrary{
					ID:   *l.SubscriptionMembershipLibraryID,
					Name: *l.SubscriptionMembershipLibraryName,
				}
			}

			membership = &models.PartialMembership{
				ID:              *l.SubscriptionMembershipID,
				LibraryID:       *l.SubscriptionMembershipLibraryID,
				Name:            *l.SubscriptionMembershipName,
				DurationDays:    *l.SubscriptionMembershipDurationDays,
				ActiveLoanLimit: *l.SubscriptionMembershipActiveLoanLimit,
				FinePerDay:      *l.SubscriptionMembershipFinePerDay,
				Library:         library,
			}
		}

		subscription = &models.PartialSubscription{
			ID:           l.SubscriptionID,
			UserID:       *l.SubscriptionUserID,
			MembershipID: *l.SubscriptionMembershipID,
			ExpiryDate:   *l.SubscriptionExpiryDate,
			User:         user,
			Membership:   membership,
		}
	}

	return &models.Loan{
		ID:             l.ID,
		LibraryBookID:  l.LibraryBookID,
		SubscriptionID: l.SubscriptionID,
		StaffID:        l.StaffID,
		LoanDate:       l.LoanDate,
		DueDate:        l.DueDate,
		ReturnDate:     l.ReturnDate,
		CreatedAt:      l.CreatedAt,
		UpdatedAt:      l.UpdatedAt,
		DeletedAt:      l.DeletedAt,

		Staff:        staff,
		LibraryBook:  libraryBook,
		Subscription: subscription,
	}
}
