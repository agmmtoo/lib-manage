package http

import (
	"context"
	"time"

	"github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

// Servicer is implemented by core/service.
type Servicer interface {
	ListUsers(ctx context.Context, input ListUsersRequest) (*ListUsersResponse, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	CreateUser(ctx context.Context, input CreateUserRequest) (*User, error)

	// ListBooks(ctx context.Context, input ListBooksRequest) (*ListBooksResponse, error)
	// GetBookByID(ctx context.Context, id int) (*models.Book, error)
	// CreateBook(ctx context.Context, input CreateBookRequest) (*models.Book, error)

	ListLibraryBooks(ctx context.Context, input ListLibraryBooksRequest) (*ListBooksResponse, error)
	GetLibraryBookByID(ctx context.Context, id int) (*models.LibraryBook, error)
	CreateBook(ctx context.Context, input CreateBookRequest) (*models.LibraryBook, error)

	ListLibraries(ctx context.Context, input ListLibrariesRequest) (*ListLibrariesResponse, error)
	GetLibraryByID(ctx context.Context, id int) (*Library, error)
	CreateLibrary(ctx context.Context, input CreateLibraryRequest) (*Library, error)
	// AssignLibraryStaff(ctx context.Context, input AssignLibraryStaffRequest) (*LibraryStaff, error)
	// RegisterLibraryBook(ctx context.Context, input RegisterLibraryBookRequest) (*LibraryBook, error)
	// RegisterLibraryBookBatch(ctx context.Context, input RegisterLibraryBookBatchRequest) (*RegisterLibraryBookBatchResponse, error)

	ListLoans(ctx context.Context, input ListLoansRequest) (*ListLoansResponse, error)
	GetLoanByID(ctx context.Context, id int) (*Loan, error)
	CreateLoan(ctx context.Context, input CreateLoanRequest) (*models.Loan, error)

	ListStaffs(ctx context.Context, input ListStaffsRequest) (*ListStaffsResponse, error)
	GetStaffByID(ctx context.Context, id int) (*Staff, error)
	CreateStaff(ctx context.Context, input CreateStaffRequest) (*Staff, error)

	ListLibrarySettings(ctx context.Context, input ListSettingsRequest) (*ListSettingsResponse, error)
	UpdateLibrarySettings(ctx context.Context, input UpdateSettingsRequest) ([]*Setting, error)

	ListMemberships(ctx context.Context, input ListMembershipsRequest) (*ListMembershipsResponse, error)

	ListSubscriptions(ctx context.Context, input ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)
	CreateSubscription(ctx context.Context, input CreateSubscriptionRequest) (*models.Subscription, error)

	// GetUsersByBookName(ctx context.Context, name string) ([]*libraryapp.User, error)

	Ping(ctx context.Context) (string, error)
	GetStats(ctx context.Context) (*Stats, error)
}

type ListResponse[T any] struct {
	Data  []*T `json:"data"`
	Total int  `json:"total"`
}

type ListUsersRequest struct {
	IDs      []int
	Limit    int
	Skip     int
	Username string
}

type ListUsersResponse = ListResponse[models.User]

type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListLibraryBooksRequest struct {
	IDs        []int
	Limit      int
	Skip       int
	Title      string
	Author     string
	LibraryIDs []int
}

type ListBooksResponse = ListResponse[models.LibraryBook]

// type Book struct {
// 	ID        int        `json:"id"`
// 	Title     string     `json:"title"`
// 	Author    string     `json:"author"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	DeletedAt *time.Time `json:"deleted_at,omitempty"`
// }

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Library struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ListLibrariesRequest struct {
	IDs   []int
	Limit int
	Skip  int
	Name  string
}

type ListLibrariesResponse = ListResponse[models.Library]

type CreateLibraryRequest struct {
	Name string `json:"name"`
}

type Loan struct {
	ID         int        `json:"id"`
	BookID     int        `json:"book_id"`
	UserID     int        `json:"user_id"`
	LibraryID  int        `json:"library_id"`
	StaffID    int        `json:"staff_id"`
	LoanDate   time.Time  `json:"loan_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`

	User    *User        `json:"user,omitempty"`
	Book    *LibraryBook `json:"book,omitempty"`
	Staff   *Staff       `json:"staff,omitempty"`
	Library *Library     `json:"library,omitempty"`
}

type ListLoansRequest struct {
	IDs            []int
	Limit          int
	Skip           int
	Active         bool
	UserIDs        []int
	LibraryBookIDs []int
	LibraryIDs     []int
	StaffIDs       []int
	DueDate        time.Time

	IncludeLibraryBook  bool
	IncludeSubscription bool
	IncludeStaff        bool
}

type ListLoansResponse = ListResponse[models.Loan]

type CreateLoanRequest struct {
	UserID        int        `json:"user_id"`
	LibraryID     int        `json:"library_id"`
	LibraryBookID int        `json:"library_book_id"`
	StaffID       int        `json:"staff_id"`
	LoanDate      *time.Time `json:"loan_date"`
	// NOTE: DueDate should be calculated based on the membership's duration days
	// DueDate       *time.Time `json:"due_date"`
}

// DEPRECATED
type Staff struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ListStaffsRequest struct {
	IDs        []int
	UserIDs    []int
	LibraryIDs []int
	Limit      int
	Skip       int
}

type ListStaffsResponse = ListResponse[models.Staff]

type CreateStaffRequest struct {
	UserID int `json:"user_id"`
}

type AssignLibraryStaffRequest struct {
	LibraryID int `json:"-"`
	StaffID   int `json:"staff_id"`
}

type LibraryStaff struct {
	LibraryID int `json:"library_id"`
	StaffID   int `json:"staff_id"`
}

type RegisterLibraryBookRequest struct {
	LibraryID int `json:"-"`
	BookID    int `json:"book_id"`
}
type RegisterLibraryBookBatchRequest struct {
	LibraryID      int   `json:"-"`
	BookIDs        []int `json:"book_ids"`
	SkipDuplicates bool  `json:"skip_duplicates"`
}

type LibraryBook struct {
	LibraryID int `json:"library_id"`
	BookID    int `json:"book_id"`
}

type RegisterLibraryBookBatchResponse struct {
	LibraryID      int   `json:"-"`
	SuccessBookIDs []int `json:"success_book_ids"`
}

type Setting struct {
	ID        int        `json:"id"`
	LibraryID int        `json:"-"`
	Key       string     `json:"key"`
	Value     string     `json:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ListSettingsRequest struct {
	IDs        []int  `json:"ids"`
	LibraryIDs []int  `json:"-"`
	Key        string `json:"key"`
	Skip       int    `json:"skip"`
	Limit      int    `json:"limit"`
}

type ListSettingsResponse = ListResponse[Setting]

type UpdateSettingsRequest = []struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Stats struct {
	Books     int `json:"books"`
	Libraries int `json:"libraries"`
	Users     int `json:"users"`
	Staffs    int `json:"staffs"`
}

type ListMembershipsRequest struct {
	IDs             []int  `json:"ids"`
	LibraryIDs      []int  `json:"library_ids"`
	Name            string `json:"name"`
	DurationDays    *int   `json:"duration_days,omitempty"`
	ActiveLoanLimit *int   `json:"active_loan_limit,omitempty"`
	FinePerDay      *int   `json:"fine_per_day,omitempty"`
	Skip            int    `json:"skip"`
	Limit           int    `json:"limit"`
}

type ListMembershipsResponse = ListResponse[models.Membership]

type ListSubscriptionsRequest struct {
	IDs           []int      `json:"ids"`
	UserIDs       []int      `json:"user_ids"`
	MembershipIDs []int      `json:"membership_ids"`
	ExpiryDate    *time.Time `json:"expiry_date,omitempty"`
	Skip          int        `json:"skip"`
	Limit         int        `json:"limit"`

	ExpiredBefore *time.Time `json:"expired_before,omitempty"`
	ExpiredAfter  *time.Time `json:"expired_after,omitempty"`
	LibraryIDs    []int      `json:"library_ids"`
}

type ListSubscriptionsResponse = ListResponse[models.Subscription]

type CreateSubscriptionRequest struct {
	UserID       int `json:"user_id"`
	MembershipID int `json:"membership_id"`
}
