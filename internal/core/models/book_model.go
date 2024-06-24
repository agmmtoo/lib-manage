package models

import (
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Core Book model
type Book struct {
	ID            int
	Title         string
	Author        string
	SubCategoryID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time

	SubCategory *PartialSubCategory
}

type PartialBook struct {
	ID            int
	Title         string
	Author        string
	SubCategoryID *int

	SubCategory *PartialSubCategory
}

// Core Book model
type LibraryBook struct {
	ID        int
	Code      string
	LibraryID int
	BookID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Book    *PartialBook
	Library *PartialLibrary
}

func (lb *LibraryBook) ToAPIModel() *models.LibraryBook {

	// TODO: implement method on source model
	var sc *models.PartialSubCategory
	if lb.Book.SubCategoryID != nil {
		sc = &models.PartialSubCategory{
			ID:         *lb.Book.SubCategoryID,
			CategoryID: lb.Book.SubCategory.Category.ID,
			Name:       lb.Book.SubCategory.Name,
			Category: &models.PartialCategory{
				ID:   lb.Book.SubCategory.Category.ID,
				Name: lb.Book.SubCategory.Category.Name,
			},
		}
	}

	return &models.LibraryBook{
		ID:        lb.ID,
		Code:      lb.Code,
		LibraryID: lb.LibraryID,
		BookID:    lb.BookID,
		CreatedAt: lb.CreatedAt,
		UpdatedAt: lb.UpdatedAt,
		DeletedAt: lb.DeletedAt,

		// TODO: implement method on source model
		Library: &models.PartialLibrary{
			ID:   lb.Library.ID,
			Name: lb.Library.Name,
		},

		// TODO: implement method on source model
		Book: &models.PartialBook{
			ID:            lb.Book.ID,
			Title:         lb.Book.Title,
			Author:        lb.Book.Author,
			SubCategoryID: lb.Book.SubCategoryID,
			SubCategory:   sc,
		},
	}
}

type PartialLibraryBook struct {
	ID        int
	Code      string
	LibraryID int
	BookID    int
	Book      *PartialBook
	Library   *PartialLibrary
}
