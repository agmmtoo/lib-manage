package library

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

func (s *Service) List(ctx context.Context, input http.ListLibrariesRequest) (*http.ListLibrariesResponse, error) {
	result, err := s.repo.ListLibraries(ctx, ListRequest{
		IDs:    input.IDs,
		Name:   input.Name,
		Limit:  input.Limit,
		Offset: input.Skip,
	})
	if err != nil {
		return nil, err
	}

	var libs []am.Library
	for _, l := range result.Libraries {
		libs = append(libs, l.ToAPIModel())
	}

	return &http.ListLibrariesResponse{
		Data:  libs,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*am.Library, error) {
	result, err := s.repo.GetLibraryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	lib := result.ToAPIModel()
	return &lib, nil
}

func (s *Service) Create(ctx context.Context, input http.CreateLibraryRequest) (*am.Library, error) {
	result, err := s.repo.CreateLibrary(ctx, CreateRequest{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	lib := result.ToAPIModel()
	return &lib, nil
}

func (s *Service) Update(ctx context.Context, id int, input http.UpdateLibraryRequest) (*am.Library, error) {
	var name *string
	if input.Name != "" {
		name = &input.Name
	}
	result, err := s.repo.UpdateLibrary(ctx, id, UpdateRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	lib := result.ToAPIModel()
	return &lib, nil
}

func (s *Service) Count(ctx context.Context) (int, error) {
	return s.repo.CountLibraries(ctx)
}

type Storer interface {
	ListLibraries(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetLibraryByID(ctx context.Context, id int) (*models.Library, error)
	CreateLibrary(ctx context.Context, input CreateRequest) (*models.Library, error)
	UpdateLibrary(ctx context.Context, id int, input UpdateRequest) (*models.Library, error)
	CountLibraries(ctx context.Context) (int, error)
}

type ListRequest struct {
	IDs    []int
	Name   string
	Limit  int
	Offset int
}

type ListResponse struct {
	Libraries []*models.Library
	Total     int
}

type CreateRequest struct {
	Name string
}

type UpdateRequest struct {
	Name *string
}

type CreateBatchOpt struct {
	SkipConflict bool
}

type CreateBatchResponse struct {
	SuccessBookIDs []int
}
