package libraryapp

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/book"
	"github.com/agmmtoo/lib-manage/internal/core/library"
	"github.com/agmmtoo/lib-manage/internal/core/loan"
	"github.com/agmmtoo/lib-manage/internal/core/membership"
	"github.com/agmmtoo/lib-manage/internal/core/setting"
	"github.com/agmmtoo/lib-manage/internal/core/staff"
	"github.com/agmmtoo/lib-manage/internal/core/subscription"
	"github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	am "github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Service implements http Service interface
// Service delegates the calls to underlying services
type Service struct {
	user         *user.Service
	book         *book.Service
	library      *library.Service
	loan         *loan.Service
	staff        *staff.Service
	setting      *setting.Service
	membership   *membership.Service
	subscription *subscription.Service
}

func New(
	user *user.Service,
	book *book.Service,
	library *library.Service,
	loan *loan.Service,
	staff *staff.Service,
	setting *setting.Service,
	membership *membership.Service,
	subscription *subscription.Service,
) *Service {
	return &Service{
		user:         user,
		book:         book,
		library:      library,
		loan:         loan,
		staff:        staff,
		setting:      setting,
		membership:   membership,
		subscription: subscription,
	}
}

func (s *Service) Ping(ctx context.Context) (string, error) {
	// TODO: delegate to the database
	return "Pong!", nil
}

func (s *Service) ListUsers(ctx context.Context, input http.ListUsersRequest) (*http.ListUsersResponse, error) {
	return s.user.List(ctx, input)
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*am.User, error) {
	return s.user.GetByID(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, input http.CreateUserRequest) (*am.User, error) {
	return s.user.Create(ctx, input)
}

func (s *Service) ListLibraryBooks(ctx context.Context, input http.ListLibraryBooksRequest) (*http.ListBooksResponse, error) {
	return s.book.ListLibraryBooks(ctx, input)
}

func (s *Service) GetLibraryBookByID(ctx context.Context, id int) (*am.LibraryBook, error) {
	return s.book.GetLibraryBookByID(ctx, id)
}

func (s *Service) CreateBook(ctx context.Context, input http.CreateBookRequest) (*am.LibraryBook, error) {
	return s.book.CreateBook(ctx, input)
}

func (s *Service) ListLibraries(ctx context.Context, input http.ListLibrariesRequest) (*http.ListLibrariesResponse, error) {
	return s.library.List(ctx, input)
}

func (s *Service) GetLibraryByID(ctx context.Context, id int) (*am.Library, error) {
	return s.library.GetByID(ctx, id)
}

func (s *Service) CreateLibrary(ctx context.Context, input http.CreateLibraryRequest) (*am.Library, error) {
	return s.library.Create(ctx, input)
}

func (s *Service) UpdateLibrary(ctx context.Context, id int, input http.UpdateLibraryRequest) (*am.Library, error) {
	return s.library.Update(ctx, id, input)
}

func (s *Service) ListLoans(ctx context.Context, input http.ListLoansRequest) (*http.ListLoansResponse, error) {
	return s.loan.List(ctx, input)
}

func (s *Service) GetLoanByID(ctx context.Context, id int) (*am.Loan, error) {
	return s.loan.GetByID(ctx, id)
}

func (s *Service) CreateLoan(ctx context.Context, input http.CreateLoanRequest) (*am.Loan, error) {
	return s.loan.Create(ctx, input)
}

func (s *Service) ListStaffs(ctx context.Context, input http.ListStaffsRequest) (*http.ListStaffsResponse, error) {
	return s.staff.List(ctx, input)
}

func (s *Service) GetStaffByID(ctx context.Context, id int) (*am.Staff, error) {
	return s.staff.GetByID(ctx, id)
}

func (s *Service) CreateStaff(ctx context.Context, input http.CreateStaffRequest) (*am.Staff, error) {
	return s.staff.Create(ctx, input)
}

// func (s *Service) AssignLibraryStaff(ctx context.Context, input http.AssignLibraryStaffRequest) (*am.LibraryStaff, error) {
// 	return s.library.AssignStaff(ctx, input)
// }

// func (s *Service) RegisterLibraryBook(ctx context.Context, input http.RegisterLibraryBookRequest) (*am.LibraryBook, error) {
// 	return s.library.RegisterBook(ctx, input)
// }

// func (s *Service) RegisterLibraryBookBatch(ctx context.Context, input http.RegisterLibraryBookBatchRequest) (*am.RegisterLibraryBookBatchResponse, error) {
// 	return s.library.RegisterBookBatch(ctx, input)
// }

func (s *Service) ListLibrarySettings(ctx context.Context, input http.ListSettingsRequest) (*http.ListSettingsResponse, error) {
	return s.setting.List(ctx, input)
}

func (s *Service) UpdateLibrarySettings(ctx context.Context, input http.UpdateSettingsRequest) ([]am.Setting, error) {
	return s.setting.Update(ctx, input)
}

func (s *Service) ListMemberships(ctx context.Context, input http.ListMembershipsRequest) (*http.ListMembershipsResponse, error) {
	return s.membership.List(ctx, input)
}

func (s *Service) ListSubscriptions(ctx context.Context, input http.ListSubscriptionsRequest) (*http.ListSubscriptionsResponse, error) {
	return s.subscription.List(ctx, input)
}

func (s *Service) CreateSubscription(ctx context.Context, input http.CreateSubscriptionRequest) (*am.Subscription, error) {
	return s.subscription.Create(ctx, input)
}

// func (s *Service) GetStats(ctx context.Context) (*am.Stats, error) {
// 	b, _ := s.book.Count(ctx)
// 	l, _ := s.library.Count(ctx)
// 	u, _ := s.user.Count(ctx)
// 	st, _ := s.staff.Count(ctx)

// 	return &http.Stats{
// 		Books:     b,
// 		Libraries: l,
// 		Users:     u,
// 		Staffs:    st,
// 	}, nil
// }

// func (s *Service) GetUsersByBookName(ctx context.Context, name string) ([]*libraryapp.User, error) {
// 	books, err := s.bookService.GetList(ctx, BookGetListInput{
// 		Name: name,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	bookIDs := make([]int, len(books))
// 	for i, b := range books {
// 		bookIDs[i] = b.ID
// 	}

// 	loans, err := s.loan.GetList(ctx, LoanGetListInput{
// 		BookIDs: bookIDs,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	userIDs := make([]int, len(loans))
// 	for i, l := range loans {
// 		userIDs[i] = l.UserID
// 	}

// 	users, err := s.userService.GetList(ctx, UserGetListInput{
// 		IDs: userIDs,
// 	})

// 	return users, err
// }
