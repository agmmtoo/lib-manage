package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/agmmtoo/lib-manage/internal/core/book"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListBooks(ctx context.Context, input book.ListRequest) (*book.ListResponse, error) {

	qb := &QueryBuilder{
		Table:        "book b",
		ParamCounter: 1,
		Cols:         []string{"b.id", "b.title", "b.author", "b.created_at", "b.updated_at", "b.deleted_at"},
	}
	if len(input.IDs) > 0 {
		qb.AddClause("b.id = ANY($%d)", input.IDs)
	}
	if len(input.Title) > 0 {
		qb.AddClause("b.title ILIKE $%d", fmt.Sprintf("%%%s%%", input.Title))
	}
	if len(input.Author) > 0 {
		qb.AddClause("b.author ILIKE $%d", fmt.Sprintf("%%%s%%", input.Author))
	}
	qb.AddClause("b.deleted_at IS NULL")
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)

	// FIXME: duplicate records on multiple IDs
	if len(input.LibraryIDs) > 0 {
		qb.JoinTables = append(qb.JoinTables, "JOIN library_book lb ON b.id = lb.book_id")
		qb.AddClause("lb.library_id = ANY($%d)", input.LibraryIDs)
	}

	q, args := qb.Build()
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

func (l *LibraryAppDB) CountBooks(ctx context.Context) (int, error) {
	var count int
	q := "SELECT count(id) from book;"
	row := l.db.QueryRowContext(ctx, q)
	err := row.Scan(&count)
	return count, err
}
