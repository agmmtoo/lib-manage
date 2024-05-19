package http

import (
	"context"
	"time"
)

// Servicer is implemented by core/service.
type Servicer interface {
	ListUsers(ctx context.Context, input ListUserRequest) (*ListUserResponse, error)
	GetUserByID(ctx context.Context, id int) (*User, error)

	ListBooks(ctx context.Context, input ListBookRequest) (*ListBookResponse, error)
	GetBookByID(ctx context.Context, id int) (*Book, error)

	// GetUsersByBookName(ctx context.Context, name string) ([]*libraryapp.User, error)

	Ping(ctx context.Context) (string, error)
}

type ListResponse[T any] struct {
	Data  []*T `json:"data"`
	Total int  `json:"total"`
}

type ListUserRequest struct {
	IDs      []int
	Limit    int
	Skip     int
	Name     string
	Username string
}

type ListUserResponse = ListResponse[User]

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Password string `json:"password"`
}

type ListBookRequest struct {
	IDs    []int
	Limit  int
	Skip   int
	Title  string
	Author string
}

type ListBookResponse = ListResponse[Book]

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
