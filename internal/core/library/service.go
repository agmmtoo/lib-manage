package library

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

	var libs []*am.Library
	for _, l := range result.Libraries {
		libs = append(libs, l.ToAPIModel())
	}

	return &http.ListLibrariesResponse{
		Data:  libs,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*http.Library, error) {
	result, err := s.repo.GetLibraryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &http.Library{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

func (s *Service) Create(ctx context.Context, input http.CreateLibraryRequest) (*http.Library, error) {
	result, err := s.repo.CreateLibrary(ctx, CreateRequest{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &http.Library{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
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
	return result.ToAPIModel(), nil
}

// func (s *Service) AssignStaff(ctx context.Context, input http.AssignLibraryStaffRequest) (*http.LibraryStaff, error) {
// 	result, err := s.repo.CreateLibraryStaff(ctx, CreateLibraryStaffRequest{
// 		LibraryID: input.LibraryID,
// 		StaffID:   input.StaffID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &http.LibraryStaff{
// 		LibraryID: result.LibraryID,
// 		StaffID:   result.StaffID,
// 	}, nil
// }

// func (s *Service) RegisterBook(ctx context.Context, input http.RegisterLibraryBookRequest) (*http.LibraryBook, error) {
// 	result, err := s.repo.CreateLibraryBook(ctx, CreateLibraryBookRequest{
// 		LibraryID: input.LibraryID,
// 		BookID:    input.BookID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &http.LibraryBook{
// 		LibraryID: result.LibraryID,
// 		BookID:    result.BookID,
// 	}, nil
// }

// func (s *Service) RegisterBookBatch(ctx context.Context, input http.RegisterLibraryBookBatchRequest) (*http.RegisterLibraryBookBatchResponse, error) {
// 	var lbs []libraryapp.LibraryBook
// 	for _, b := range input.BookIDs {
// 		lbs = append(lbs, libraryapp.LibraryBook{
// 			LibraryID: input.LibraryID,
// 			BookID:    b,
// 		})
// 	}

// 	result, err := s.repo.CreateLibraryBookBatch(ctx, lbs, CreateBatchOpt{
// 		SkipConflict: input.SkipDuplicates,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &http.RegisterLibraryBookBatchResponse{
// 		LibraryID:      input.LibraryID,
// 		SuccessBookIDs: result.SuccessBookIDs,
// 	}, nil
// }

func (s *Service) Count(ctx context.Context) (int, error) {
	return s.repo.CountLibraries(ctx)
}

type Storer interface {
	ListLibraries(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetLibraryByID(ctx context.Context, id int) (*libraryapp.Library, error)
	CreateLibrary(ctx context.Context, input CreateRequest) (*libraryapp.Library, error)
	UpdateLibrary(ctx context.Context, id int, input UpdateRequest) (*models.Library, error)
	// CreateLibraryStaff(ctx context.Context, input CreateLibraryStaffRequest) (*libraryapp.LibraryStaff, error)
	// CreateLibraryBook(ctx context.Context, input CreateLibraryBookRequest) (*libraryapp.LibraryBook, error)
	// CreateLibraryBookBatch(ctx context.Context, input []libraryapp.LibraryBook, opt CreateBatchOpt) (*CreateBatchResponse, error)
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

type CreateLibraryStaffRequest struct {
	LibraryID int
	StaffID   int
}

type CreateLibraryBookRequest struct {
	LibraryID int
	BookID    int
}

type CreateBatchOpt struct {
	SkipConflict bool
}

type CreateBatchResponse struct {
	SuccessBookIDs []int
}
