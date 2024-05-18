package book

import (
	"context"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, input ListRequest) (*ListResponse, error) {
	return s.repo.ListBooks(ctx, input)
}

func (s *Service) GetByID(ctx context.Context, id int) (*libraryapp.Book, error) {
	return s.repo.GetBookByID(ctx, id)
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
