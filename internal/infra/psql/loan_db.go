package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/agmmtoo/lib-manage/internal/core/loan"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
	"github.com/jackc/pgx/v5/pgconn"
)

func (l *LibraryAppDB) ListLoans(ctx context.Context, input loan.ListRequest) (*loan.ListResponse, error) {

	qb := &QueryBuilder{
		Table:        "loan l",
		ParamCounter: 1,
		Cols:         []string{"l.id", "l.book_id", "l.user_id", "l.library_id", "l.staff_id", "l.loan_date", "l.due_date", "l.return_date", "l.created_at", "l.updated_at", "l.deleted_at"},
	}

	if input.IncludeUser {
		qb.JoinTables = append(qb.JoinTables, "JOIN \"user\" u ON l.user_id = u.id")
		qb.Cols = append(qb.Cols, "u.id", "u.username", "u.created_at", "u.updated_at", "u.deleted_at")
	}

	if input.IncludeBook {
		qb.JoinTables = append(qb.JoinTables, "JOIN \"book\" b ON l.book_id = b.id")
		qb.Cols = append(qb.Cols, "b.id", "b.title", "b.author", "b.created_at", "b.updated_at", "b.deleted_at")
	}

	if input.IncludeStaff {
		qb.JoinTables = append(qb.JoinTables, "JOIN \"staff\" s ON l.staff_id = s.id")
		qb.Cols = append(qb.Cols, "s.id", "s.user_id", "s.created_at", "s.updated_at", "s.deleted_at")
	}

	if input.IncludeLibrary {
		qb.JoinTables = append(qb.JoinTables, "JOIN \"library\" lb ON l.library_id = lb.id")
		qb.Cols = append(qb.Cols, "lb.id", "lb.name", "lb.created_at", "lb.updated_at", "lb.deleted_at")
	}

	if len(input.IDs) > 0 {
		qb.AddClause("l.id = ANY($%d)", input.IDs)
	}
	if input.Active {
		qb.AddClause("l.return_date IS NULL")
	}
	if len(input.UserIDs) > 0 {
		qb.AddClause("l.user_id = ANY($%d)", input.UserIDs)
	}
	if len(input.BookIDs) > 0 {
		qb.AddClause("l.book_id = ANY($%d)", input.BookIDs)
	}
	if len(input.LibraryIDs) > 0 {
		qb.AddClause("l.library_id = ANY($%d)", input.LibraryIDs)
	}
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)

	query, params := qb.Build()
	rows, err := l.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing loans", err)
	}

	defer rows.Close()

	var loans []*libraryapp.Loan
	for rows.Next() {
		var l libraryapp.Loan
		dests := []interface{}{&l.ID, &l.BookID, &l.UserID, &l.LibraryID, &l.StaffID, &l.LoanDate, &l.DueDate, &l.ReturnDate, &l.CreatedAt, &l.UpdatedAt, &l.DeletedAt}
		if input.IncludeUser {
			l.User = &libraryapp.User{}
			dests = append(dests, &l.User.ID, &l.User.Username, &l.User.CreatedAt, &l.User.UpdatedAt, &l.User.DeletedAt)
		}
		if input.IncludeBook {
			l.Book = &libraryapp.Book{}
			dests = append(dests, &l.Book.ID, &l.Book.Title, &l.Book.Author, &l.Book.CreatedAt, &l.Book.UpdatedAt, &l.Book.DeletedAt)
		}
		if input.IncludeStaff {
			l.Staff = &libraryapp.Staff{}
			dests = append(dests, &l.Staff.ID, &l.Staff.UserID, &l.Staff.CreatedAt, &l.Staff.UpdatedAt, &l.Staff.DeletedAt)
		}
		if input.IncludeLibrary {
			l.Library = &libraryapp.Library{}
			dests = append(dests, &l.Library.ID, &l.Library.Name, &l.Library.CreatedAt, &l.Library.UpdatedAt, &l.Library.DeletedAt)
		}
		err := rows.Scan(dests...)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error scanning loan", err)
		}
		loans = append(loans, &l)
	}

	if err := rows.Err(); err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing loans", err)
	}

	return &loan.ListResponse{
		Loans: loans,
	}, nil
}

func (l *LibraryAppDB) GetLoanByID(ctx context.Context, id int) (*libraryapp.Loan, error) {
	q := "SELECT id, book_id, user_id, library_id, staff_id, loan_date, due_date, return_date, created_at, updated_at, deleted_at FROM loan WHERE id = $1;"
	args := []any{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lo libraryapp.Loan
	err := row.Scan(&lo.ID, &lo.BookID, &lo.UserID, &lo.LibraryID, &lo.StaffID, &lo.LoanDate, &lo.DueDate, &lo.ReturnDate, &lo.CreatedAt, &lo.UpdatedAt, &lo.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "loan not found", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error getting loan", err)
	}

	return &lo, nil
}

func (l *LibraryAppDB) CreateLoan(ctx context.Context, input loan.CreateRequest) (*libraryapp.Loan, error) {
	q := "INSERT INTO loan (user_id, book_id, library_id, staff_id, loan_date, due_date, return_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, user_id, book_id, library_id, staff_id, loan_date, due_date, return_date, created_at, updated_at, deleted_at;"
	args := []any{input.UserID, input.BookID, input.LibraryID, input.StaffID, input.LoanDate, input.DueDate, input.ReturnDate}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lo libraryapp.Loan
	err := row.Scan(&lo.ID, &lo.UserID, &lo.BookID, &lo.LibraryID, &lo.StaffID, &lo.LoanDate, &lo.DueDate, &lo.ReturnDate, &lo.CreatedAt, &lo.UpdatedAt, &lo.DeletedAt)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Println(pgErr.Detail)
			if pgErr.ConstraintName == "loan_user_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "user not found", err)
			}
			if pgErr.ConstraintName == "loan_library_id_staff_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "library or staff not found", err)
			}
			if pgErr.ConstraintName == "loan_library_id_book_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "library or book not found", err)
			}
			if pgErr.ConstraintName == "loan_unique_active" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBDuplicate, "book already loaned", err)
			}
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error creating loan", err)
	}

	return &lo, nil
}
