package loan

import (
	"context"
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, input http.ListLoansRequest) (*http.ListLoansResponse, error) {
	result, err := s.repo.ListLoans(ctx, ListRequest{
		IDs:    input.IDs,
		Limit:  input.Limit,
		Offset: input.Skip,
		// UserID: input.UserID,
		// BookID: input.BookID,
	})
	if err != nil {
		return nil, err
	}
	var loans []*http.Loan
	for _, l := range result.Loans {
		loans = append(loans, &http.Loan{
			ID:         l.ID,
			UserID:     l.UserID,
			BookID:     l.BookID,
			LibraryID:  l.LibraryID,
			StaffID:    l.StaffID,
			LoanDate:   l.LoanDate,
			DueDate:    l.DueDate,
			ReturnDate: l.ReturnDate,
			CreatedAt:  l.CreatedAt,
			UpdatedAt:  l.UpdatedAt,
			DeletedAt:  l.DeletedAt,
		})
	}

	return &http.ListLoansResponse{
		Data:  loans,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*http.Loan, error) {
	result, err := s.repo.GetLoanByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &http.Loan{
		ID:         result.ID,
		UserID:     result.UserID,
		BookID:     result.BookID,
		LibraryID:  result.LibraryID,
		StaffID:    result.StaffID,
		LoanDate:   result.LoanDate,
		DueDate:    result.DueDate,
		ReturnDate: result.ReturnDate,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
		DeletedAt:  result.DeletedAt,
	}, nil
}

func (s *Service) Create(ctx context.Context, input http.CreateLoanRequest) (*http.Loan, error) {

	// TODO:
	// - check book is available
	// - check user is not over limit
	// - set due date based on library rules

	loanDate := time.Now()
	if input.LoanDate != nil {
		loanDate = *input.LoanDate
	}

	dueDate := time.Now().AddDate(0, 0, 7)
	if input.DueDate != nil {
		dueDate = *input.DueDate
	}

	result, err := s.repo.CreateLoan(ctx, CreateRequest{
		UserID:    input.UserID,
		BookID:    input.BookID,
		LibraryID: input.LibraryID,
		StaffID:   input.StaffID,
		LoanDate:  loanDate,
		DueDate:   dueDate,
		// ReturnDate: input.ReturnDate,
	})
	if err != nil {
		return nil, err
	}
	return &http.Loan{
		ID:         result.ID,
		UserID:     result.UserID,
		BookID:     result.BookID,
		LibraryID:  result.LibraryID,
		StaffID:    result.StaffID,
		LoanDate:   result.LoanDate,
		DueDate:    result.DueDate,
		ReturnDate: result.ReturnDate,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
		DeletedAt:  result.DeletedAt,
	}, nil
}

type Storer interface {
	ListLoans(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetLoanByID(ctx context.Context, id int) (*libraryapp.Loan, error)
	CreateLoan(ctx context.Context, input CreateRequest) (*libraryapp.Loan, error)
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

type CreateRequest struct {
	UserID     int
	BookID     int
	LibraryID  int
	StaffID    int
	LoanDate   time.Time
	DueDate    time.Time
	ReturnDate *time.Time
}
