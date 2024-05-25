package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/core/loan"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
	"github.com/lib/pq"
)

func (l *LibraryAppDB) ListLoans(ctx context.Context, input loan.ListRequest) (*loan.ListResponse, error) {
	// q := "SELECT id, book_id, user_id, library_id, staff_id, loan_date, due_date, return_date, created_at, updated_at, deleted_at FROM loan"
	// args := []any{}

	baseQuery := "SELECT * FROM loan WHERE 1=1"
	var params []interface{}
	var conditions []string
	var paramCounter = 1

	if len(input.IDs) > 0 {
		conditions = append(conditions, fmt.Sprintf("id = ANY($%d)", paramCounter))
		params = append(params, pq.Array(input.IDs))
		paramCounter++
	}
	if input.Active {
		conditions = append(conditions, fmt.Sprintf("return_date IS NULL"))
	}
	if len(input.UserIDs) > 0 {
		conditions = append(conditions, fmt.Sprintf("user_id = ANY($%d)", paramCounter))
		params = append(params, pq.Array(input.UserIDs))
		paramCounter++
	}
	if len(input.BookIDs) > 0 {
		conditions = append(conditions, fmt.Sprintf("book_id = ANY($%d)", paramCounter))
		params = append(params, pq.Array(input.BookIDs))
		paramCounter++
	}
	if len(input.LibraryIDs) > 0 {
		conditions = append(conditions, fmt.Sprintf("library_id = ANY($%d)", paramCounter))
		params = append(params, pq.Array(input.LibraryIDs))
		paramCounter++
	}

	finalQuery := baseQuery + " " + strings.Join(conditions, " AND ")

	if input.Limit > 0 {
		finalQuery += fmt.Sprintf(" LIMIT $%d", paramCounter)
		params = append(params, input.Limit)
		paramCounter++
	}
	if input.Offset > 0 {
		finalQuery += fmt.Sprintf(" OFFSET $%d", paramCounter)
		params = append(params, input.Offset)
		paramCounter++
	}

	fmt.Println(finalQuery, params)
	rows, err := l.db.QueryContext(ctx, finalQuery, params...)
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
		if pqErr, ok := err.(*pq.Error); ok {
			fmt.Println(pqErr.Detail)
			if pqErr.Constraint == "loan_user_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "user not found", err)
			}
			if pqErr.Constraint == "loan_library_id_staff_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "library or staff not found", err)
			}
			if pqErr.Constraint == "loan_library_id_book_id_fkey" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "library or book not found", err)
			}
			if pqErr.Constraint == "loan_unique_active" {
				return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBDuplicate, "book already loaned", err)
			}
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error creating loan", err)
	}

	return &lo, nil
}
