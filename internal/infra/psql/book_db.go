package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/agmmtoo/lib-manage/internal/core"
	"github.com/agmmtoo/lib-manage/internal/core/book"
	cm "github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/infra/psql/models"
)

func (l *LibraryAppDB) ListLibraryBooks(ctx context.Context, input book.ListRequest) (*book.ListResponse, error) {

	qb := &QueryBuilder{
		Table:        "libraries_books lb",
		ParamCounter: 1,
		Cols: []string{
			"lb.id", "lb.code", "lb.library_id", "lb.book_id", "lb.created_at", "lb.updated_at", "lb.deleted_at",
			"l.id", "l.name",
			"b.id", "b.title", "b.author", "b.sub_category_id",
			"sc.id", "sc.category_id", "sc.name",
			"c.id", "c.name",
		},
	}
	if len(input.IDs) > 0 {
		qb.AddClause("lb.id = ANY($%d)", input.IDs)
	}
	if len(input.Title) > 0 {
		qb.AddClause("b.title ILIKE $%d", fmt.Sprintf("%%%s%%", input.Title))
	}
	if len(input.Author) > 0 {
		qb.AddClause("b.author ILIKE $%d", fmt.Sprintf("%%%s%%", input.Author))
	}

	// Join with "libraries"
	qb.JoinTables = append(qb.JoinTables, "JOIN libraries l ON lb.library_id = l.id")

	// Join with "books"
	qb.JoinTables = append(qb.JoinTables, "JOIN books b ON lb.book_id = b.id")

	// Join with "sub_categories"
	qb.JoinTables = append(qb.JoinTables, "LEFT JOIN sub_categories sc ON b.sub_category_id = sc.id")

	// Join with "categories"
	qb.JoinTables = append(qb.JoinTables, "LEFT JOIN categories c ON sc.category_id = c.id")

	qb.AddClause("b.deleted_at IS NULL")
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)

	if len(input.LibraryIDs) > 0 {
		qb.AddClause("lb.library_id = ANY($%d)", input.LibraryIDs)
	}

	q, args := qb.Build()
	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing books", err)
	}

	defer rows.Close()

	var books []cm.LibraryBook

	for rows.Next() {
		var b models.LibraryBook
		err := rows.Scan(
			&b.ID, &b.Code, &b.LibraryID, &b.BookID, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt,
			&b.LibraryID, &b.LibraryName,
			&b.BookID, &b.BookTitle, &b.BookAuthor, &b.BookSubCategoryID,
			&b.BookSubCategoryID, &b.BookSubCategoryCategoryID, &b.BookSubCategoryName,
			&b.BookSubCategoryCategoryID, &b.BookSubCategoryCategoryName,
		)
		if err != nil {
			return nil, core.NewCoreError(core.ErrCodeDBScan, "error scanning book", err)
		}

		books = append(books, b.ToCoreModel())
	}

	if err := rows.Err(); err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing books", err)
	}

	var count int
	// TODO: implement count query
	// qc, p := qb.BuildCount()
	// row := l.db.QueryRowContext(ctx, qc, p)
	// err = row.Scan(&count)
	// if err != nil {
	// 	fmt.Printf("Count scan error: %v\n", err)
	// }

	return &book.ListResponse{
		Books: books,
		Total: count,
	}, nil
}

func (l *LibraryAppDB) GetLibraryBookByID(ctx context.Context, id int) (*cm.LibraryBook, error) {
	qb := &QueryBuilder{
		Table:        "libraries_books lb",
		ParamCounter: 1,
		Cols: []string{
			"lb.id", "lb.code", "lb.library_id", "lb.book_id", "lb.created_at", "lb.updated_at", "lb.deleted_at",
			"l.id", "l.name",
			"b.id", "b.title", "b.author", "b.sub_category_id",
			"sc.id", "sc.category_id", "sc.name",
			"c.id", "c.name",
		},
	}

	// Join with "libraries"
	qb.JoinTables = append(qb.JoinTables, "JOIN libraries l ON lb.library_id = l.id")
	// Join with "books"
	qb.JoinTables = append(qb.JoinTables, "JOIN books b ON lb.book_id = b.id")
	// Join with "sub_categories"
	qb.JoinTables = append(qb.JoinTables, "LEFT JOIN sub_categories sc ON b.sub_category_id = sc.id")
	// Join with "categories"
	qb.JoinTables = append(qb.JoinTables, "LEFT JOIN categories c ON sc.category_id = c.id")
	qb.AddClause("b.deleted_at IS NULL")

	qb.AddClause("lb.id = $%d", id)

	q, args := qb.Build()

	row := l.db.QueryRowContext(ctx, q, args...)

	var b models.LibraryBook
	err := row.Scan(
		&b.ID, &b.Code, &b.LibraryID, &b.BookID, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt,
		&b.LibraryID, &b.LibraryName,
		&b.BookID, &b.BookTitle, &b.BookAuthor, &b.BookSubCategoryID,
		&b.BookSubCategoryID, &b.BookSubCategoryCategoryID, &b.BookSubCategoryName,
		&b.BookSubCategoryCategoryID, &b.BookSubCategoryCategoryName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "book not found", err)
		}
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error getting book", err)
	}

	book := b.ToCoreModel()
	return &book, nil
}

func (l *LibraryAppDB) CreateBook(ctx context.Context, input book.CreateRequest) (*cm.LibraryBook, error) {
	q := "INSERT INTO book (title, author) VALUES ($1, $2) returning id, title, author, created_at, updated_at, deleted_at;"
	args := []any{input.Title, input.Arthor}

	row := l.db.QueryRowContext(ctx, q, args...)

	var b models.LibraryBook
	err := row.Scan(&b.ID, &b.BookTitle, &b.BookAuthor, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt)
	if err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBScan, "error creating book", err)
	}
	book := b.ToCoreModel()
	return &book, nil
}

func (l *LibraryAppDB) CountBooks(ctx context.Context) (int, error) {
	var count int
	q := "SELECT count(id) from book;"
	row := l.db.QueryRowContext(ctx, q)
	err := row.Scan(&count)
	return count, err
}
