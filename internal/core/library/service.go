package library

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
	return s.repo.ListLibraries(ctx, input)
}

func (s *Service) GetByID(ctx context.Context, id int) (*libraryapp.Library, error) {
	return s.repo.GetLibraryByID(ctx, id)
}

type Storer interface {
	ListLibraries(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetLibraryByID(ctx context.Context, id int) (*libraryapp.Library, error)
}

type ListRequest struct {
	IDs    []int
	Name   string
	Limit  int
	Offset int
}

type ListResponse struct {
	Libraries []*libraryapp.Library
	Total     int
}
