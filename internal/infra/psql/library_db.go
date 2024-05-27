package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/core/library"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
	"github.com/lib/pq"
)

func (l *LibraryAppDB) ListLibraries(ctx context.Context, input library.ListRequest) (*library.ListResponse, error) {
	// q := "SELECT id, name, created_at, updated_at, deleted_at FROM library;"
	// args := []any{}
	qb := &QueryBuilder{
		Table:        "library",
		ParamCounter: 1,
	}
	if len(input.IDs) > 0 {
		qb.AddClause("id = ANY($%d)", pq.Array(input.IDs))
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
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing libraries", err)
	}

	defer rows.Close()

	var libraries []*libraryapp.Library
	for rows.Next() {
		var lib libraryapp.Library
		err := rows.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error scanning library", err)
		}
		libraries = append(libraries, &lib)
	}

	if err := rows.Err(); err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing libraries", err)
	}

	return &library.ListResponse{
		Libraries: libraries,
	}, nil
}

func (l *LibraryAppDB) GetLibraryByID(ctx context.Context, id int) (*libraryapp.Library, error) {
	q := "SELECT id, name, created_at, updated_at, deleted_at FROM library WHERE id = $1;"
	args := []any{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lib libraryapp.Library
	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "library not found", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error getting library", err)
	}

	return &lib, nil
}

func (l *LibraryAppDB) CreateLibrary(ctx context.Context, input library.CreateRequest) (*libraryapp.Library, error) {
	q := "INSERT INTO library (name) VALUES ($1) RETURNING id, name, created_at, updated_at, deleted_at;"
	args := []any{input.Name}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lib libraryapp.Library
	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error creating library", err)
	}

	// Create default setting for the library
	qs := "INSERT INTO setting (library_id, key, value) VALUES ($1, $2, $3), ($1, $4, $5), ($1, $6, $7);"
	_, err = l.db.ExecContext(ctx, qs, lib.ID,
		config.SETTING_KEY_MAX_LOAN_PER_USER, config.SETTING_DEFAULT_MAX_LOAN_PER_USER,
		config.SETTING_KEY_LOAN_PERIOD, config.SETTING_DEFAULT_LOAN_PERIOD,
		config.SETTING_KEY_FINE_PER_DAY, config.SETTING_DEFAULT_FINE_PER_DAY,
	)

	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBExec, "error creating library settings", err)
	}

	return &lib, nil
}

func (l *LibraryAppDB) CreateLibraryStaff(ctx context.Context, input library.CreateLibraryStaffRequest) (*libraryapp.LibraryStaff, error) {
	q := "INSERT INTO library_staff (library_id, staff_id) VALUES ($1, $2) RETURNING library_id, staff_id, created_at, updated_at, deleted_at;"
	args := []any{input.LibraryID, input.StaffID}

	row := l.db.QueryRowContext(ctx, q, args...)

	var libStaff libraryapp.LibraryStaff
	err := row.Scan(&libStaff.LibraryID, &libStaff.StaffID, &libStaff.CreatedAt, &libStaff.UpdatedAt, &libStaff.DeletedAt)
	if err != nil {
		if strings.Contains(err.Error(), "library_staff_library_id_fkey") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "library not found", err)
		}
		if strings.Contains(err.Error(), "library_staff_staff_id_fkey") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "staff not found", err)
		}
		if strings.Contains(err.Error(), "library_staff_pkey") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBDuplicate, "staff is already assigned", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error creating library staff", err)
	}

	return &libStaff, nil
}

func (l *LibraryAppDB) CreateLibraryBook(ctx context.Context, input library.CreateLibraryBookRequest) (*libraryapp.LibraryBook, error) {
	q := "INSERT INTO library_book (library_id, book_id) VALUES ($1, $2) RETURNING library_id, book_id, created_at, updated_at, deleted_at;"
	args := []any{input.LibraryID, input.BookID}

	row := l.db.QueryRowContext(ctx, q, args...)

	var libBook libraryapp.LibraryBook
	err := row.Scan(&libBook.LibraryID, &libBook.BookID, &libBook.CreatedAt, &libBook.UpdatedAt, &libBook.DeletedAt)
	if err != nil {
		if strings.Contains(err.Error(), "library_book_library_id_fkey") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "library not found", err)
		}
		if strings.Contains(err.Error(), "library_book_book_id_fkey") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "book not found", err)
		}
		if strings.Contains(err.Error(), "library_book_pkey") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBDuplicate, "book is already registered", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error creating library book", err)
	}

	return &libBook, nil
}

func (l *LibraryAppDB) CreateLibraryBookBatch(ctx context.Context, input []libraryapp.LibraryBook, opt library.CreateBatchOpt) (*library.CreateBatchResponse, error) {
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
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Constraint == "library_book_pkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBDuplicate, pqErr.Detail, err)
			}
			if pqErr.Constraint == "library_book_book_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, pqErr.Detail, err)
			}
			if pqErr.Constraint == "library_book_library_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, pqErr.Detail, err)
			}
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBExec, "error creating library book batch", err)
	}

	defer rows.Close()

	for rows.Next() {
		var bookID int
		err := rows.Scan(&bookID)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error scanning book id", err)
		}
		successBookIDs = append(successBookIDs, bookID)
	}

	return &library.CreateBatchResponse{
		SuccessBookIDs: successBookIDs,
	}, nil
}
