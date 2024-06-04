package loan

import (
	"context"
	"strconv"
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, input http.ListLoansRequest) (*http.ListLoansResponse, error) {
	result, err := s.repo.ListLoans(ctx, ListRequest{
		IDs:            input.IDs,
		Limit:          input.Limit,
		Offset:         input.Skip,
		Active:         input.Active,
		UserIDs:        input.UserIDs,
		BookIDs:        input.BookIDs,
		LibraryIDs:     input.LibraryIDs,
		StaffIDs:       input.StaffIDs,
		IncludeUser:    input.IncludeUser,
		IncludeBook:    input.IncludeBook,
		IncludeLibrary: input.IncludeLibrary,
		IncludeStaff:   input.IncludeStaff,
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
			// FIXME: nil pointer dereference
			User: &http.User{
				ID:        l.User.ID,
				Username:  l.User.Username,
				CreatedAt: l.User.CreatedAt,
				UpdatedAt: l.User.UpdatedAt,
				DeletedAt: l.User.DeletedAt,
			},
			// FIXME: nil pointer dereference
			Book: &http.Book{
				ID:        l.Book.ID,
				Title:     l.Book.Title,
				Author:    l.Book.Author,
				CreatedAt: l.Book.CreatedAt,
				UpdatedAt: l.Book.UpdatedAt,
				DeletedAt: l.Book.DeletedAt,
			},
			// FIXME: nil pointer dereference
			Staff: &http.Staff{
				ID:        l.Staff.ID,
				UserID:    l.Staff.UserID,
				CreatedAt: l.Staff.CreatedAt,
				UpdatedAt: l.Staff.UpdatedAt,
				DeletedAt: l.Staff.DeletedAt,
			},
			// FIXME: nil pointer dereference
			Library: &http.Library{
				ID:        l.Library.ID,
				Name:      l.Library.Name,
				CreatedAt: l.Library.CreatedAt,
				UpdatedAt: l.Library.UpdatedAt,
				DeletedAt: l.Library.DeletedAt,
			},
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

	loanDate := time.Now()
	if input.LoanDate != nil {
		loanDate = *input.LoanDate
	}

	// check user's loan limit setting
	settingLimit, err := s.repo.GetSettingValue(ctx, input.LibraryID, config.SETTING_KEY_MAX_LOAN_PER_USER)
	if err != nil {
		return nil, err
	}
	limit, err := strconv.Atoi(settingLimit)
	if err != nil {
		return nil, err
	}
	activeLoans, err := s.repo.ListLoans(ctx, ListRequest{
		UserIDs:    []int{input.UserID},
		LibraryIDs: []int{input.LibraryID},
		Active:     true,
		Limit:      limit,
	})
	if err != nil {
		return nil, err
	}
	if len(activeLoans.Loans) >= limit {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeForbidden, "user has reached loan limit", nil)
	}

	// set due date based on loan period setting
	settingPeriod, err := s.repo.GetSettingValue(ctx, input.LibraryID, config.SETTING_KEY_LOAN_PERIOD)
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(settingPeriod)
	if err != nil {
		return nil, err
	}
	dueDate := time.Now().AddDate(0, 0, day)
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

	GetSettingValue(ctx context.Context, libraryID int, key string) (string, error)
}

type ListRequest struct {
	IDs        []int
	Active     bool
	UserIDs    []int
	BookIDs    []int
	LibraryIDs []int
	StaffIDs   []int
	Limit      int
	Offset     int

	IncludeUser    bool
	IncludeBook    bool
	IncludeStaff   bool
	IncludeLibrary bool
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
