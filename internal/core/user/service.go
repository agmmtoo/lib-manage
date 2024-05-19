package user

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

func (s *Service) List(ctx context.Context, input http.ListUserRequest) (*http.ListUserResponse, error) {
	result, err := s.repo.ListUsers(ctx, ListRequest{
		IDs:      input.IDs,
		Name:     input.Name,
		Username: input.Username,
		Limit:    input.Limit,
		Offset:   input.Skip,
	})

	if err != nil {
		return nil, err
	}

	var users []*http.User
	for _, u := range result.Users {
		users = append(users, &http.User{
			ID:        u.ID,
			Username:  u.Username,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		})
	}

	return &http.ListUserResponse{
		Data:  users,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*http.User, error) {
	result, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &http.User{
		ID:        result.ID,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

// Storer provides access to the user storage
// Implemented by the database layer, `internal/infra/psql`
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
