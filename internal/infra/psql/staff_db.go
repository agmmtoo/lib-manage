package psql

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/staff"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListStaffs(ctx context.Context, input staff.ListRequest) (*staff.ListResponse, error) {
	q := "SELECT id, user_id, created_at, updated_at, deleted_at FROM staff;"
	args := []any{}
	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var staffs []*libraryapp.Staff
	for rows.Next() {
		var s libraryapp.Staff
		err := rows.Scan(&s.ID, &s.UserID, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &staff.ListResponse{
		Staffs: staffs,
	}, nil
}

func (l *LibraryAppDB) GetStaffByID(ctx context.Context, id int) (*libraryapp.Staff, error) {
	q := "SELECT id, user_id, created_at, updated_at, deleted_at FROM staff WHERE id = $1;"
	args := []any{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var s libraryapp.Staff
	err := row.Scan(&s.ID, &s.UserID, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (l *LibraryAppDB) CreateStaff(ctx context.Context, input staff.CreateRequest) (*libraryapp.Staff, error) {
	q := "INSERT INTO staff (user_id) VALUES ($1) RETURNING id, user_id, created_at, updated_at, deleted_at;"
	args := []any{input.UserID}

	row := l.db.QueryRowContext(ctx, q, args...)

	var s libraryapp.Staff
	err := row.Scan(&s.ID, &s.UserID, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
