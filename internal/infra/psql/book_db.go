package psql

import (
	"context"
	"database/sql"

	"github.com/agmmtoo/lib-manage/internal/core/book"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListBooks(ctx context.Context, input book.ListRequest) (*book.ListResponse, error) {
	q := "SELECT id, title, author, created_at, updated_at, deleted_at FROM book;"
	args := []any{}
	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing books", err)
	}

	defer rows.Close()

	var books []*libraryapp.Book
	for rows.Next() {
		var b libraryapp.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error scanning book", err)
		}
		books = append(books, &b)
	}

	if err := rows.Err(); err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing books", err)
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
		if err == sql.ErrNoRows {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "book not found", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error getting book", err)
	}

	return &b, nil
}

func (l *LibraryAppDB) CreateBook(ctx context.Context, input book.CreateRequest) (*libraryapp.Book, error) {
	q := "INSERT INTO book (title, author) VALUES ($1, $2) returning id, title, author, created_at, updated_at, deleted_at;"
	args := []any{input.Title, input.Arthor}

	row := l.db.QueryRowContext(ctx, q, args...)

	var b libraryapp.Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt)
	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error creating book", err)
	}
	return &b, nil
}
