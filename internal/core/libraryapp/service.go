package libraryapp

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/book"
	"github.com/agmmtoo/lib-manage/internal/core/library"
	"github.com/agmmtoo/lib-manage/internal/core/loan"
	"github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Service struct {
	user    *user.Service
	book    *book.Service
	library *library.Service
	loan    *loan.Service
}

func New(
	user *user.Service,
	book *book.Service,
	library *library.Service,
	loan *loan.Service,
) *Service {
	return &Service{
		user:    user,
		book:    book,
		library: library,
		loan:    loan,
	}
}

func (s *Service) ListUsers(ctx context.Context, input user.ListRequest) (*user.ListResponse, error) {
	return s.user.List(ctx, input)
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*libraryapp.User, error) {
	return s.user.GetByID(ctx, id)
}

func (s *Service) ListBooks(ctx context.Context, input book.ListRequest) (*book.ListResponse, error) {
	return s.book.List(ctx, input)
}

func (s *Service) GetBookByID(ctx context.Context, id int) (*libraryapp.Book, error) {
	return s.book.GetByID(ctx, id)
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
