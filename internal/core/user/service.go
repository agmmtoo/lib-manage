package user

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
	return s.repo.ListUsers(ctx, input)
}

func (s *Service) GetByID(ctx context.Context, id int) (*libraryapp.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

// Storer provides access to the user storage.
type Storer interface {
	ListUsers(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetUserByID(ctx context.Context, id int) (*libraryapp.User, error)
}

type ListRequest struct {
	IDs      []int
	Name     string
	Username string
	Limit    int
	Offset   int
}

type ListResponse struct {
	Users []*libraryapp.User
	Total int
}
