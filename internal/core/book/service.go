package book

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	am "github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListLibraryBooks(ctx context.Context, input http.ListLibraryBooksRequest) (*http.ListBooksResponse, error) {
	result, err := s.repo.ListLibraryBooks(ctx, ListRequest{
		IDs:        input.IDs,
		Title:      input.Title,
		Author:     input.Author,
		Limit:      input.Limit,
		Offset:     input.Skip,
		LibraryIDs: input.LibraryIDs,
	})
	if err != nil {
		return nil, err
	}

	var books []*am.LibraryBook
	for _, b := range result.Books {
		books = append(books, b.ToAPIModel())
	}

	return &http.ListBooksResponse{
		Data:  books,
		Total: result.Total,
	}, nil
}

func (s *Service) GetLibraryBookByID(ctx context.Context, id int) (*am.LibraryBook, error) {
	b, err := s.repo.GetLibraryBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return b.ToAPIModel(), nil
}

func (s *Service) CreateBook(ctx context.Context, input http.CreateBookRequest) (*am.LibraryBook, error) {
	result, err := s.repo.CreateBook(ctx, CreateRequest{
		Title:  input.Title,
		Arthor: input.Author,
	})
	if err != nil {
		return nil, err
	}
	return result.ToAPIModel(), nil
}

func (s *Service) Count(ctx context.Context) (int, error) {
	return s.repo.CountBooks(ctx)
}

type Storer interface {
	ListLibraryBooks(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetLibraryBookByID(ctx context.Context, id int) (*models.LibraryBook, error)
	CreateBook(ctx context.Context, input CreateRequest) (*models.LibraryBook, error)
	CountBooks(ctx context.Context) (int, error)
}

type ListRequest struct {
	IDs        []int
	Title      string
	Author     string
	Limit      int
	Offset     int
	LibraryIDs []int
}

type ListResponse struct {
	Books []*models.LibraryBook
	Total int
}

type CreateRequest struct {
	Title         string
	Arthor        string
	SubCategoryID *int
	Code          string
}
