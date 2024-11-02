package loan

import (
	"context"
	"fmt"
	"time"

	"github.com/agmmtoo/lib-manage/internal/core"
	"github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/core/subscription"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	am "github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, input http.ListLoansRequest) (*http.ListLoansResponse, error) {
	result, err := s.repo.ListLoans(ctx, ListRequest{
		IDs:                 input.IDs,
		Limit:               input.Limit,
		Offset:              input.Skip,
		Active:              input.Active,
		UserIDs:             input.UserIDs,
		LibraryBookIDs:      input.LibraryBookIDs,
		LibraryIDs:          input.LibraryIDs,
		StaffIDs:            input.StaffIDs,
		IncludeLibraryBook:  input.IncludeLibraryBook,
		IncludeSubscription: input.IncludeSubscription,
		IncludeStaff:        input.IncludeStaff,
	})
	if err != nil {
		return nil, err
	}
	var loans []am.Loan
	for _, l := range result.Loans {
		loans = append(loans, l.ToAPIModel())
	}

	return &http.ListLoansResponse{
		Data:  loans,
		Total: result.Total,
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*am.Loan, error) {
	result, err := s.repo.GetLoanByID(ctx, id)
	if err != nil {
		return nil, err
	}
	loan := result.ToAPIModel()
	return &loan, nil
}

func (s *Service) Create(ctx context.Context, input http.CreateLoanRequest) (*am.Loan, error) {

	loanDate := time.Now()
	if input.LoanDate != nil {
		loanDate = *input.LoanDate
	}

	// get library id from book
	book, err := s.repo.GetLibraryBookByID(ctx, input.LibraryBookID)
	if err != nil {
		return nil, err
	}
	// insanity check
	if input.LibraryID != book.LibraryID {
		return nil, core.NewCoreError(core.ErrCodeBadRequest, "book not found in library", nil)
	}

	// check book is available
	activeLoans, err := s.repo.ListLoans(ctx, ListRequest{
		LibraryBookIDs: []int{input.LibraryBookID},
		Active:         true,
	})
	if err != nil {
		return nil, err
	}
	if len(activeLoans.Loans) > 0 {
		// DEBUG
		fmt.Println("Book's active loans: ", time.Until(activeLoans.Loans[0].DueDate))
		return nil, core.NewCoreError(core.ErrCodeForbidden, "book is not available", nil)
	}

	// get library id from staff
	staff, err := s.repo.GetStaffByID(ctx, input.StaffID)
	if err != nil {
		return nil, err
	}
	// insanity check
	if staff.LibraryID != input.LibraryID {
		return nil, core.NewCoreError(core.ErrCodeBadRequest, "staff is not right for library", nil)
	}

	// check user has active subscription
	subs, err := s.repo.ListSubscriptions(ctx, subscription.ListRequest{
		UserIDs:    []int{input.UserID},
		LibraryIDs: []int{input.LibraryID},
		OrderBy: []struct {
			Col string
			Dir string
		}{{Col: "created_at", Dir: "desc"}},
	})
	if err != nil {
		return nil, err
	}
	if len(subs.Subscriptions) == 0 {
		return nil, core.NewCoreError(core.ErrCodeForbidden, "user has no active subscription for library", nil)
	}

	// NOTE: only the latest subscription is considered
	// check subscription is active
	sub := subs.Subscriptions[0]

	if sub.ExpiryDate.Before(loanDate) {
		return nil, core.NewCoreError(core.ErrCodeForbidden, "user subscription has expired", nil)
	}

	// get membership
	mbs, err := s.repo.GetMembershipByID(ctx, sub.MembershipID)
	if err != nil {
		return nil, err
	}

	// get user's active loans
	ual, err := s.repo.ListLoans(ctx, ListRequest{
		UserIDs:             []int{input.UserID},
		LibraryIDs:          []int{input.LibraryID},
		IncludeSubscription: true,
		Active:              true,
	})
	if err != nil {
		return nil, err
	}

	// check user has reached loan limit
	if len(ual.Loans) >= mbs.ActiveLoanLimit {
		return nil, core.NewCoreError(core.ErrCodeForbidden, fmt.Sprintf("user has reached loan limit: %d for membership: %d", mbs.ActiveLoanLimit, mbs.ID), nil)
	}

	// set due date based on membership
	dueDate := time.Now().AddDate(0, 0, mbs.DurationDays)
	// if input.DueDate != nil {
	// 	dueDate = *input.DueDate
	// }

	result, err := s.repo.CreateLoan(ctx, CreateRequest{
		SubscriptionID: sub.ID,
		LibraryBookID:  input.LibraryBookID,
		StaffID:        input.StaffID,
		LoanDate:       loanDate,
		DueDate:        dueDate,
		// ReturnDate: input.ReturnDate,
	})
	if err != nil {
		return nil, err
	}
	return &am.Loan{
		ID:             result.ID,
		LibraryBookID:  result.LibraryBookID,
		SubscriptionID: result.SubscriptionID,
		StaffID:        result.StaffID,
		LoanDate:       result.LoanDate,
		DueDate:        result.DueDate,
		ReturnDate:     result.ReturnDate,
		CreatedAt:      result.CreatedAt,
		UpdatedAt:      result.UpdatedAt,
		DeletedAt:      result.DeletedAt,
	}, nil
}

type Storer interface {
	ListLoans(ctx context.Context, input ListRequest) (*ListResponse, error)
	GetLoanByID(ctx context.Context, id int) (*models.Loan, error)
	CreateLoan(ctx context.Context, input CreateRequest) (*models.Loan, error)

	// implemented by book service
	GetLibraryBookByID(ctx context.Context, id int) (*models.LibraryBook, error)
	// implemented by staff service
	GetStaffByID(ctx context.Context, id int) (*models.Staff, error)
	// implemented by subscription service
	ListSubscriptions(ctx context.Context, input subscription.ListRequest) (*subscription.ListResponse, error)
	// implemented by membership service
	GetMembershipByID(ctx context.Context, id int) (*models.Membership, error)
}

type ListRequest struct {
	IDs            []int
	Active         bool
	UserIDs        []int
	LibraryBookIDs []int
	LibraryIDs     []int
	StaffIDs       []int
	ExpiryDate     *time.Time
	Limit          int
	Offset         int

	IncludeLibraryBook  bool
	IncludeSubscription bool
	IncludeStaff        bool
}

type ListResponse struct {
	Loans []*models.Loan
	Total int
}

type CreateRequest struct {
	LibraryBookID  int
	SubscriptionID int
	StaffID        int
	LoanDate       time.Time
	DueDate        time.Time
	ReturnDate     *time.Time
}
