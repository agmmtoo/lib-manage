package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Core Loan model
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

	Staff        *PartialStaff
	LibraryBook  *PartialLibraryBook
	Subscription *PartialSubscription
}

func (l Loan) ToAPIModel() models.Loan {

	var lb *models.PartialLibraryBook
	if l.LibraryBook != nil {

		var b *models.PartialBook
		if l.LibraryBook.Book != nil {

			var sc *models.PartialSubCategory
			if l.LibraryBook.Book.SubCategory != nil {

				var c *models.PartialCategory
				if l.LibraryBook.Book.SubCategory.Category != nil {
					c = &models.PartialCategory{
						ID:   l.LibraryBook.Book.SubCategory.Category.ID,
						Name: l.LibraryBook.Book.SubCategory.Category.Name,
					}
				}
				sc = &models.PartialSubCategory{
					ID:       l.LibraryBook.Book.SubCategory.ID,
					Name:     l.LibraryBook.Book.SubCategory.Name,
					Category: c,
				}
			}
			b = &models.PartialBook{
				ID:            l.LibraryBook.Book.ID,
				Title:         l.LibraryBook.Book.Title,
				Author:        l.LibraryBook.Book.Author,
				SubCategoryID: l.LibraryBook.Book.SubCategoryID,
				SubCategory:   sc,
			}
		}

		var lib *models.PartialLibrary
		if l.LibraryBook.Library != nil {
			lib = &models.PartialLibrary{
				ID:   l.LibraryBook.Library.ID,
				Name: l.LibraryBook.Library.Name,
			}
		}

		lb = &models.PartialLibraryBook{
			ID:        l.LibraryBook.ID,
			LibraryID: l.LibraryBook.LibraryID,
			BookID:    l.LibraryBook.BookID,
			Code:      l.LibraryBook.Code,
			Book:      b,
			Library:   lib,
		}
	}

	var s *models.PartialStaff
	if l.Staff != nil {
		var u *models.PartialUser
		if l.Staff.User != nil {
			u = &models.PartialUser{
				ID:       l.Staff.User.ID,
				Username: l.Staff.User.Username,
			}
		}

		var lib *models.PartialLibrary
		if l.Staff.Library != nil {
			lib = &models.PartialLibrary{
				ID:   l.Staff.Library.ID,
				Name: l.Staff.Library.Name,
			}
		}

		s = &models.PartialStaff{
			ID:        l.Staff.ID,
			UserID:    l.Staff.UserID,
			LibraryID: l.Staff.LibraryID,
			User:      u,
			Library:   lib,
		}
	}

	var sub *models.PartialSubscription
	if l.Subscription != nil {

		var u *models.PartialUser
		if l.Subscription.User != nil {
			u = &models.PartialUser{
				ID:       l.Subscription.User.ID,
				Username: l.Subscription.User.Username,
			}
		}

		var m *models.PartialMembership
		if l.Subscription.Membership != nil {

			var lib *models.PartialLibrary
			if l.Subscription.Membership.Library != nil {
				lib = &models.PartialLibrary{
					ID:   l.Subscription.Membership.Library.ID,
					Name: l.Subscription.Membership.Library.Name,
				}
			}

			m = &models.PartialMembership{
				ID:              l.Subscription.Membership.ID,
				Name:            l.Subscription.Membership.Name,
				LibraryID:       l.Subscription.Membership.LibraryID,
				DurationDays:    l.Subscription.Membership.DurationDays,
				ActiveLoanLimit: l.Subscription.Membership.ActiveLoanLimit,
				FinePerDay:      l.Subscription.Membership.FinePerDay,
				Library:         lib,
			}
		}

		sub = &models.PartialSubscription{
			ID:           l.Subscription.ID,
			UserID:       l.Subscription.UserID,
			MembershipID: l.Subscription.MembershipID,
			ExpiryDate:   l.Subscription.ExpiryDate,
			User:         u,
			Membership:   m,
		}
	}

	return models.Loan{
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

		LibraryBook:  lb,
		Staff:        s,
		Subscription: sub,
	}
}
