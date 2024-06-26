package membership

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

func (s *Service) List(ctx context.Context, input http.ListMembershipsRequest) (*http.ListMembershipsResponse, error) {
	result, err := s.repo.ListMemberships(ctx, ListRequest{
		IDs:             input.IDs,
		LibraryIDs:      input.LibraryIDs,
		Name:            input.Name,
		DurationDays:    input.DurationDays,
		ActiveLoanLimit: input.ActiveLoanLimit,
		FinePerDay:      input.FinePerDay,
		Limit:           input.Limit,
		Offset:          input.Skip,
	})
	if err != nil {
		return nil, err
	}

	var memberships []*am.Membership
	for _, m := range result.Memberships {
		memberships = append(memberships, m.ToAPIModel())
	}

	return &http.ListMembershipsResponse{
		Data:  memberships,
		Total: result.Total,
	}, nil
}

type Storer interface {
	ListMemberships(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetMembershipByID(ctx context.Context, id int) (*models.Membership, error)
}

type ListRequest struct {
	IDs             []int
	LibraryIDs      []int
	Name            string
	DurationDays    *int
	ActiveLoanLimit *int
	FinePerDay      *int
	Limit           int
	Offset          int
}

type ListResponse struct {
	Memberships []*models.Membership
	Total       int
}
