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
		Table:        "loan",
		ParamCounter: 1,
	}

	if len(input.IDs) > 0 {
		qb.AddClause("id = ANY($%d)", input.IDs)
	}
	if input.Active {
		qb.AddClause("return_date IS NULL")
	}
	if len(input.UserIDs) > 0 {
		qb.AddClause("user_id = ANY($%d)", input.UserIDs)
	}
	if len(input.BookIDs) > 0 {
		qb.AddClause("book_id = ANY($%d)", input.BookIDs)
	}
	if len(input.LibraryIDs) > 0 {
		qb.AddClause("library_id = ANY($%d)", input.LibraryIDs)
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
		err := rows.Scan(&l.ID, &l.BookID, &l.UserID, &l.LibraryID, &l.StaffID, &l.LoanDate, &l.DueDate, &l.ReturnDate, &l.CreatedAt, &l.UpdatedAt, &l.DeletedAt)
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
