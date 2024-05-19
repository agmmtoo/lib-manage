package http

import (
	"context"
	"time"
)

// Servicer is implemented by core/service.
type Servicer interface {
	ListUsers(ctx context.Context, input ListUsersRequest) (*ListUsersResponse, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	CreateUser(ctx context.Context, input CreateUserRequest) (*User, error)

	ListBooks(ctx context.Context, input ListBooksRequest) (*ListBooksResponse, error)
	GetBookByID(ctx context.Context, id int) (*Book, error)
	CreateBook(ctx context.Context, input CreateBookRequest) (*Book, error)

	ListLibraries(ctx context.Context, input ListLibrariesRequest) (*ListLibrariesResponse, error)
	GetLibraryByID(ctx context.Context, id int) (*Library, error)
	CreateLibrary(ctx context.Context, input CreateLibraryRequest) (*Library, error)

	ListLoans(ctx context.Context, input ListLoansRequest) (*ListLoansResponse, error)
	GetLoanByID(ctx context.Context, id int) (*Loan, error)
	CreateLoan(ctx context.Context, input CreateLoanRequest) (*Loan, error)

	ListStaffs(ctx context.Context, input ListStaffsRequest) (*ListStaffsResponse, error)
	GetStaffByID(ctx context.Context, id int) (*Staff, error)
	CreateStaff(ctx context.Context, input CreateStaffRequest) (*Staff, error)

	// GetUsersByBookName(ctx context.Context, name string) ([]*libraryapp.User, error)

	Ping(ctx context.Context) (string, error)
}

type ListResponse[T any] struct {
	Data  []*T `json:"data"`
	Total int  `json:"total"`
}

type ListUsersRequest struct {
	IDs      []int
	Limit    int
	Skip     int
	Name     string
	Username string
}

type ListUsersResponse = ListResponse[User]

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

type ListBooksRequest struct {
	IDs    []int
	Limit  int
	Skip   int
	Title  string
	Author string
}

type ListBooksResponse = ListResponse[Book]

type Book struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

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

type ListLibrariesResponse = ListResponse[Library]

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
}

type ListLoansRequest struct {
	IDs     []int
	Limit   int
	Skip    int
	DueDate time.Time
}

type ListLoansResponse = ListResponse[Loan]

type CreateLoanRequest struct {
	ID        int        `json:"id"`
	BookID    int        `json:"book_id"`
	UserID    int        `json:"user_id"`
	LibraryID int        `json:"library_id"`
	StaffID   int        `json:"staff_id"`
	LoanDate  *time.Time `json:"loan_date"`
	DueDate   *time.Time `json:"due_date"`
}

type Staff struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ListStaffsRequest struct {
	IDs   []int
	Limit int
	Skip  int
}

type ListStaffsResponse = ListResponse[Staff]

type CreateStaffRequest struct {
	UserID int `json:"user_id"`
}
