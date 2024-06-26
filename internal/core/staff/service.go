package staff

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	am "github.com/agmmtoo/lib-manage/internal/infra/http/models"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, input http.ListStaffsRequest) (*http.ListStaffsResponse, error) {
	result, err := s.repo.ListStaffs(ctx, ListRequest{
		IDs:    input.IDs,
		Limit:  input.Limit,
		Offset: input.Skip,
		// Name: input.Name,
		UserIDs:    input.UserIDs,
		LibraryIDs: input.LibraryIDs,
	})
	if err != nil {
		return nil, err
	}
	var staffs []*am.Staff
	for _, s := range result.Staffs {
		staffs = append(staffs, s.ToAPIModel())
	}
	return &http.ListStaffsResponse{
		Data:  staffs,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*http.Staff, error) {
	result, err := s.repo.GetStaffByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &http.Staff{
		ID:        result.ID,
		UserID:    result.UserID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

func (s *Service) Create(ctx context.Context, input http.CreateStaffRequest) (*http.Staff, error) {
	result, err := s.repo.CreateStaff(ctx, CreateRequest{
		UserID: input.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &http.Staff{
		ID:        result.ID,
		UserID:    result.UserID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

func (s *Service) Count(ctx context.Context) (int, error) {
	return s.repo.CountStaffs(ctx)
}

type Storer interface {
	ListStaffs(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetStaffByID(ctx context.Context, id int) (*libraryapp.Staff, error)
	CreateStaff(ctx context.Context, input CreateRequest) (*libraryapp.Staff, error)
	CountStaffs(ctx context.Context) (int, error)
}

type ListRequest struct {
	IDs        []int
	UserIDs    []int
	LibraryIDs []int
	Limit      int
	Offset     int
}

type ListResponse struct {
	Staffs []*models.Staff
	Total  int
}

type CreateRequest struct {
	UserID int
}
