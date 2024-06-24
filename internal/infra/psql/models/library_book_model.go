package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
)

// LibraryBook represents LibraryBook model, joined Library and Book
type LibraryBook struct {
	ID        int
	LibraryID int
	BookID    int
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// joined fields from Library model
	LibraryName *string

	// joined fields from Book model
	BookTitle                   *string
	BookAuthor                  *string
	BookSubCategoryID           *int
	BookSubCategoryName         *string
	BookSubCategoryCategoryID   *int
	BookSubCategoryCategoryName *string
}

func (lb *LibraryBook) ToCoreModel() *models.LibraryBook {

	// TODO: implement method on source model
	var sc *models.PartialSubCategory
	if lb.BookSubCategoryID != nil {
		sc = &models.PartialSubCategory{
			ID:         *lb.BookSubCategoryID,
			CategoryID: *lb.BookSubCategoryCategoryID,
			Name:       *lb.BookSubCategoryName,
			Category: &models.PartialCategory{
				ID:   *lb.BookSubCategoryCategoryID,
				Name: *lb.BookSubCategoryCategoryName,
			},
		}
	}

	return &models.LibraryBook{
		ID:        lb.ID,
		LibraryID: lb.LibraryID,
		BookID:    lb.BookID,
		Code:      lb.Code,
		CreatedAt: lb.CreatedAt,
		UpdatedAt: lb.UpdatedAt,
		DeletedAt: lb.DeletedAt,

		// TODO: implement method on source model
		Library: &models.PartialLibrary{
			ID:   lb.LibraryID,
			Name: *lb.LibraryName,
		},

		// TODO: implement method on source model
		Book: &models.PartialBook{
			ID:            lb.BookID,
			Title:         *lb.BookTitle,
			Author:        *lb.BookAuthor,
			SubCategoryID: lb.BookSubCategoryID,
			SubCategory:   sc,
		},
	}
}
