package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/agmmtoo/lib-manage/internal/core"
	"github.com/agmmtoo/lib-manage/internal/core/library"
	cm "github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/infra/psql/models"
	"github.com/jackc/pgx/v5/pgconn"
)

func (l *LibraryAppDB) ListLibraries(ctx context.Context, input library.ListRequest) (*library.ListResponse, error) {
	qb := &QueryBuilder{
		Table:        "libraries",
		ParamCounter: 1,
	}
	if len(input.IDs) > 0 {
		qb.AddClause("id = ANY($%d)", input.IDs)
	}
	if len(input.Name) > 0 {
		qb.AddClause("name ILIKE $%d", fmt.Sprintf("%%%s%%", input.Name))
	}
	qb.AddClause("deleted_at IS NULL")
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)
	q, args := qb.Build()

	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing libraries", err)
	}

	defer rows.Close()

	var libraries []*cm.Library
	for rows.Next() {
		var lib models.Library
		err := rows.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
		if err != nil {
			return nil, core.NewCoreError(core.ErrCodeDBScan, "error scanning library", err)
		}
		libraries = append(libraries, lib.ToCoreModel())
	}

	if err := rows.Err(); err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing libraries", err)
	}

	return &library.ListResponse{
		Libraries: libraries,
	}, nil
}

func (l *LibraryAppDB) GetLibraryByID(ctx context.Context, id int) (*cm.Library, error) {
	q := "SELECT id, name, created_at, updated_at, deleted_at FROM libraries WHERE id = $1;"
	args := []any{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lib models.Library
	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "library not found", err)
		}
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error getting library", err)
	}

	return lib.ToCoreModel(), nil
}

func (l *LibraryAppDB) CreateLibrary(ctx context.Context, input library.CreateRequest) (*cm.Library, error) {
	q := "INSERT INTO libraries (name) VALUES ($1) RETURNING id, name, created_at, updated_at, deleted_at;"
	args := []any{input.Name}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lib models.Library
	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
	if err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBScan, "error creating library", err)
	}

	return lib.ToCoreModel(), nil
}

func (l *LibraryAppDB) UpdateLibrary(ctx context.Context, id int, input library.UpdateRequest) (*cm.Library, error) {
	uvm := make(map[string]any)
	if input.Name != nil {
		uvm["name"] = *input.Name
	}
	var (
		uc   string
		args = []any{
			id,
		}
	)

	for k, v := range uvm {
		uc += fmt.Sprintf("%s = $%d, ", k, len(args)+1)
		args = append(args, v)
	}

	q := fmt.Sprintf("UPDATE libraries SET %s updated_at = now() WHERE id = $1 RETURNING id, name, created_at, updated_at, deleted_at;", uc)

	row := l.db.QueryRowContext(ctx, q, args...)

	var lib models.Library
	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
	if err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBScan, "error updating library", err)
	}
	return lib.ToCoreModel(), nil
}

// FIXME: Duplicate code
// func (l *LibraryAppDB) CreateLibraryBook(ctx context.Context, input library.CreateLibraryBookRequest) (*cm.LibraryBook, error) {
// 	q := "INSERT INTO library_book (library_id, book_id) VALUES ($1, $2) RETURNING library_id, book_id, created_at, updated_at, deleted_at;"
// 	args := []any{input.LibraryID, input.BookID}

// 	row := l.db.QueryRowContext(ctx, q, args...)

// 	var libBook cm.LibraryBook
// 	err := row.Scan(&libBook.LibraryID, &libBook.BookID, &libBook.CreatedAt, &libBook.UpdatedAt, &libBook.DeletedAt)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "library_book_library_id_fkey") {
// 			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "library not found", err)
// 		}
// 		if strings.Contains(err.Error(), "library_book_book_id_fkey") {
// 			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "book not found", err)
// 		}
// 		if strings.Contains(err.Error(), "library_book_pkey") {
// 			return nil, core.NewCoreError(core.ErrCodeDBDuplicate, "book is already registered", err)
// 		}
// 		return nil, core.NewCoreError(core.ErrCodeDBScan, "error creating library book", err)
// 	}

// 	return &libBook, nil
// }

func (l *LibraryAppDB) CreateLibraryBookBatch(ctx context.Context, input []cm.LibraryBook, opt library.CreateBatchOpt) (*library.CreateBatchResponse, error) {
	var successBookIDs []int

	q := "INSERT INTO library_book (library_id, book_id) VALUES "
	vals := []any{}

	for i, lb := range input {
		q += fmt.Sprintf("($%d, $%d),", i*2+1, i*2+2)
		vals = append(vals, lb.LibraryID, lb.BookID)
	}

	// Remove the last comma
	q = q[:len(q)-1]

	if opt.SkipConflict {
		q += " ON CONFLICT DO NOTHING"
	}

	q += " RETURNING book_id;"

	rows, err := l.db.QueryContext(ctx, q, vals...)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.ConstraintName == "library_book_pkey" {
				return nil, core.NewCoreError(core.ErrCodeDBDuplicate, pgErr.Detail, err)
			}
			if pgErr.ConstraintName == "library_book_book_id_fkey" {
				return nil, core.NewCoreError(core.ErrCodeDBNotFound, pgErr.Detail, err)
			}
			if pgErr.ConstraintName == "library_book_library_id_fkey" {
				return nil, core.NewCoreError(core.ErrCodeDBNotFound, pgErr.Detail, err)
			}
		}
		return nil, core.NewCoreError(core.ErrCodeDBExec, "error creating library book batch", err)
	}

	defer rows.Close()

	for rows.Next() {
		var bookID int
		err := rows.Scan(&bookID)
		if err != nil {
			return nil, core.NewCoreError(core.ErrCodeDBScan, "error scanning book id", err)
		}
		successBookIDs = append(successBookIDs, bookID)
	}

	return &library.CreateBatchResponse{
		SuccessBookIDs: successBookIDs,
	}, nil
}

func (l *LibraryAppDB) CountLibraries(ctx context.Context) (int, error) {
	var count int
	q := "SELECT count(id) from library;"
	row := l.db.QueryRowContext(ctx, q)
	err := row.Scan(&count)
	return count, err
}
