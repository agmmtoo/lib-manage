package psql

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/loan"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListLoans(ctx context.Context, input loan.ListRequest) (*loan.ListResponse, error) {
	q := "SELECT id, user_id, book_id, created_at, updated_at, deleted_at FROM loan;"
	args := []interface{}{}
	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var loans []*libraryapp.Loan
	for rows.Next() {
		var l libraryapp.Loan
		err := rows.Scan(&l.ID, &l.UserID, &l.BookID, &l.CreatedAt, &l.UpdatedAt, &l.DeletedAt)
		if err != nil {
			return nil, err
		}
		loans = append(loans, &l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (l *LibraryAppDB) GetLoanByID(ctx context.Context, id int) (*libraryapp.Loan, error) {
	q := "SELECT id, user_id, book_id, created_at, updated_at, deleted_at FROM loan WHERE id = $1;"
	args := []interface{}{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lo libraryapp.Loan
	err := row.Scan(&lo.ID, &lo.UserID, &lo.BookID, &lo.CreatedAt, &lo.UpdatedAt, &lo.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &lo, nil
}
