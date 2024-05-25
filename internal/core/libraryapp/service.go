package libraryapp

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/book"
	"github.com/agmmtoo/lib-manage/internal/core/library"
	"github.com/agmmtoo/lib-manage/internal/core/loan"
	"github.com/agmmtoo/lib-manage/internal/core/staff"
	"github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
)

// Service implements http Service interface
// Service delegates the calls to underlying services
type Service struct {
	user    *user.Service
	book    *book.Service
	library *library.Service
	loan    *loan.Service
	staff   *staff.Service
}

func New(
	user *user.Service,
	book *book.Service,
	library *library.Service,
	loan *loan.Service,
	staff *staff.Service,
) *Service {
	return &Service{
		user:    user,
		book:    book,
		library: library,
		loan:    loan,
		staff:   staff,
	}
}

func (s *Service) Ping(ctx context.Context) (string, error) {
	// TODO: delegate to the database
	return "Pong!", nil
}

func (s *Service) ListUsers(ctx context.Context, input http.ListUsersRequest) (*http.ListUsersResponse, error) {
	return s.user.List(ctx, input)
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*http.User, error) {
	return s.user.GetByID(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, input http.CreateUserRequest) (*http.User, error) {
	return s.user.Create(ctx, input)
}

func (s *Service) ListBooks(ctx context.Context, input http.ListBooksRequest) (*http.ListBooksResponse, error) {
	return s.book.List(ctx, input)
}

func (s *Service) GetBookByID(ctx context.Context, id int) (*http.Book, error) {
	return s.book.GetByID(ctx, id)
}

func (s *Service) CreateBook(ctx context.Context, input http.CreateBookRequest) (*http.Book, error) {
	return s.book.Create(ctx, input)
}

func (s *Service) ListLibraries(ctx context.Context, input http.ListLibrariesRequest) (*http.ListLibrariesResponse, error) {
	return s.library.List(ctx, input)
}

func (s *Service) GetLibraryByID(ctx context.Context, id int) (*http.Library, error) {
	return s.library.GetByID(ctx, id)
}

func (s *Service) CreateLibrary(ctx context.Context, input http.CreateLibraryRequest) (*http.Library, error) {
	return s.library.Create(ctx, input)
}

func (s *Service) ListLoans(ctx context.Context, input http.ListLoansRequest) (*http.ListLoansResponse, error) {
	return s.loan.List(ctx, input)
}

func (s *Service) GetLoanByID(ctx context.Context, id int) (*http.Loan, error) {
	return s.loan.GetByID(ctx, id)
}

func (s *Service) CreateLoan(ctx context.Context, input http.CreateLoanRequest) (*http.Loan, error) {
	return s.loan.Create(ctx, input)
}

func (s *Service) ListStaffs(ctx context.Context, input http.ListStaffsRequest) (*http.ListStaffsResponse, error) {
	return s.staff.List(ctx, input)
}

func (s *Service) GetStaffByID(ctx context.Context, id int) (*http.Staff, error) {
	return s.staff.GetByID(ctx, id)
}

func (s *Service) CreateStaff(ctx context.Context, input http.CreateStaffRequest) (*http.Staff, error) {
	return s.staff.Create(ctx, input)
}

func (s *Service) AssignLibraryStaff(ctx context.Context, input http.AssignLibraryStaffRequest) (*http.LibraryStaff, error) {
	return s.library.AssignStaff(ctx, input)
}

func (s *Service) RegisterLibraryBook(ctx context.Context, input http.RegisterLibraryBookRequest) (*http.LibraryBook, error) {
	return s.library.RegisterBook(ctx, input)
}

func (s *Service) RegisterLibraryBookBatch(ctx context.Context, input http.RegisterLibraryBookBatchRequest) (*http.RegisterLibraryBookBatchResponse, error) {
	return s.library.RegisterBookBatch(ctx, input)
}

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
