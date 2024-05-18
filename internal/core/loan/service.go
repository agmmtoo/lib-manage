package loan

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
	return s.repo.ListLoans(ctx, input)
}

func (s *Service) GetByID(ctx context.Context, id int) (*libraryapp.Loan, error) {
	return s.repo.GetLoanByID(ctx, id)
}

type Storer interface {
	ListLoans(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetLoanByID(ctx context.Context, id int) (*libraryapp.Loan, error)
}

type ListRequest struct {
	IDs    []int
	UserID int
	BookID int
	Limit  int
	Offset int
}

type ListResponse struct {
	Loans []*libraryapp.Loan
	Total int
}
