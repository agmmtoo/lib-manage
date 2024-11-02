package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

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

func (s *Service) List(ctx context.Context, input http.ListUsersRequest) (*http.ListUsersResponse, error) {
	result, err := s.repo.ListUsers(ctx, ListRequest{
		IDs:      input.IDs,
		Username: input.Username,
		Limit:    input.Limit,
		Offset:   input.Skip,
	})

	if err != nil {
		return nil, err
	}

	var users []am.User
	for _, u := range result.Users {
		users = append(users, u.ToAPIModel())
	}

	return &http.ListUsersResponse{
		Data:  users,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*am.User, error) {
	result, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &am.User{
		ID:        result.ID,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

func (s *Service) Create(ctx context.Context, input http.CreateUserRequest) (*am.User, error) {
	// hash the password
	h := sha256.New()
	h.Write([]byte(input.Password))
	b := h.Sum(nil)
	password := hex.EncodeToString(b)

	result, err := s.repo.CreateUser(ctx, CreateRequest{
		Username: input.Username,
		Password: string(password),
	})

	if err != nil {
		return nil, err
	}

	return &am.User{
		ID:        result.ID,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

func (s *Service) Count(ctx context.Context) (int, error) {
	return s.repo.CountUsers(ctx)
}

// Storer provides access to the user storage
// Implemented by the database layer, `internal/infra/psql`
type Storer interface {
	ListUsers(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, input CreateRequest) (*models.User, error)
	CountUsers(ctx context.Context) (int, error)
}

type ListRequest struct {
	IDs      []int
	Name     string
	Username string
	Limit    int
	Offset   int
}

type ListResponse struct {
	Users []*models.User
	Total int
}

type CreateRequest struct {
	Username string
	Password string
}
