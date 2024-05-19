package psql

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/book"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListBooks(ctx context.Context, input book.ListRequest) (*book.ListResponse, error) {
	q := "SELECT id, title, author, created_at, updated_at, deleted_at FROM book;"
	args := []interface{}{}
	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*libraryapp.Book
	for rows.Next() {
		var b libraryapp.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &book.ListResponse{
		Books: books,
	}, nil
}

func (l *LibraryAppDB) GetBookByID(ctx context.Context, id int) (*libraryapp.Book, error) {
	q := "SELECT id, title, author, created_at, updated_at, deleted_at FROM book WHERE id = $1;"
	args := []any{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var b libraryapp.Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
