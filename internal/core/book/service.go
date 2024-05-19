package book

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/infra/http"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, input http.ListBookRequest) (*http.ListBookResponse, error) {
	result, err := s.repo.ListBooks(ctx, ListRequest{
		IDs:    input.IDs,
		Title:  input.Title,
		Author: input.Author,
		Limit:  input.Limit,
		Offset: input.Skip,
	})
	if err != nil {
		return nil, err
	}

	var books []*http.Book
	for _, b := range result.Books {
		books = append(books, &http.Book{
			ID:        b.ID,
			Title:     b.Title,
			Author:    b.Author,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
			DeletedAt: b.DeletedAt,
		})
	}

	return &http.ListBookResponse{
		Data:  books,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*http.Book, error) {
	result, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &http.Book{
		ID:        result.ID,
		Title:     result.Title,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

type Storer interface {
	ListBooks(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetBookByID(ctx context.Context, id int) (*libraryapp.Book, error)
}

type ListRequest struct {
	IDs    []int
	Title  string
	Author string
	Limit  int
	Offset int
}

type ListResponse struct {
	Books []*libraryapp.Book
	Total int
}
