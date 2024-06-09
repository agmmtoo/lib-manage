package psql

import (
	"context"
	"database/sql"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/core/staff"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListStaffs(ctx context.Context, input staff.ListRequest) (*staff.ListResponse, error) {
	qb := &QueryBuilder{
		Table:        "staff s",
		ParamCounter: 1,
		Cols:         []string{"s.id", "s.user_id", "s.created_at", "s.updated_at", "s.deleted_at"},
	}
	if len(input.IDs) > 0 {
		qb.AddClause("s.id = ANY($%d)", input.IDs)
	}
	if len(input.UserIDs) > 0 {
		qb.AddClause("s.user_id = ANY($%d)", input.UserIDs)
	}
	if len(input.LibraryIDs) > 0 {
		qb.JoinTables = append(qb.JoinTables, "JOIN library_staff ls ON s.id = ls.staff_id")
		qb.AddClause("ls.library_id = ANY($%d)", input.LibraryIDs)
	}
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)
	q, args := qb.Build()

	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing staffs", err)
	}

	defer rows.Close()

	var staffs []*libraryapp.Staff
	for rows.Next() {
		var s libraryapp.Staff
		err := rows.Scan(&s.ID, &s.UserID, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error scanning staff", err)
		}
		staffs = append(staffs, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing staffs", err)
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
		if err == sql.ErrNoRows {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "staff not found", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error getting staff", err)
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
		if strings.Contains(err.Error(), "staff_user_id_fkey") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "staff user doesn't exist", err)
		}
		if strings.Contains(err.Error(), "staff_user_id_key") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBDuplicate, "user is already a staff", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error creating staff", err)
	}

	return &s, nil
}
