package psql

import (
	"context"
	"database/sql"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/core"
	cm "github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/core/staff"
	"github.com/agmmtoo/lib-manage/internal/infra/psql/models"
)

func (l *LibraryAppDB) ListStaffs(ctx context.Context, input staff.ListRequest) (*staff.ListResponse, error) {
	qb := &QueryBuilder{
		Table:        "staffs s",
		ParamCounter: 1,
		Cols: []string{
			"s.id", "s.user_id", "s.library_id", "s.created_at", "s.updated_at", "s.deleted_at",
			"l.id", "l.name",
			"u.id", "u.username",
		},
	}
	if len(input.IDs) > 0 {
		qb.AddClause("s.id = ANY($%d)", input.IDs)
	}
	if len(input.UserIDs) > 0 {
		qb.AddClause("s.user_id = ANY($%d)", input.UserIDs)
	}

	// Join with "libraries"
	qb.JoinTables = append(qb.JoinTables, "JOIN libraries l ON s.library_id = l.id")

	// Join with "users"
	qb.JoinTables = append(qb.JoinTables, "JOIN users u ON s.user_id = u.id")

	if len(input.LibraryIDs) > 0 {
		qb.AddClause("s.library_id = ANY($%d)", input.LibraryIDs)
	}

	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)
	q, args := qb.Build()

	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing staffs", err)
	}

	defer rows.Close()

	var staffs []*cm.Staff
	for rows.Next() {
		var s models.Staff
		err := rows.Scan(
			&s.ID, &s.UserID, &s.LibraryID, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt,
			&s.LibraryID, &s.LibraryName,
			&s.UserID, &s.UserUsername,
		)
		if err != nil {
			return nil, core.NewCoreError(core.ErrCodeDBScan, "error scanning staff", err)
		}
		staffs = append(staffs, s.ToCoreModel())
	}

	if err := rows.Err(); err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing staffs", err)
	}

	return &staff.ListResponse{
		Staffs: staffs,
	}, nil
}

func (l *LibraryAppDB) GetStaffByID(ctx context.Context, id int) (*cm.Staff, error) {
	q := "SELECT id, user_id, library_id, created_at, updated_at, deleted_at FROM staffs WHERE id = $1;"
	args := []any{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var s models.Staff
	err := row.Scan(&s.ID, &s.UserID, &s.LibraryID, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	// TODO: handle pg error (old code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "staff not found", err)
		}
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error getting staff", err)
	}

	return s.ToCoreModel(), nil
}

func (l *LibraryAppDB) CreateStaff(ctx context.Context, input staff.CreateRequest) (*cm.Staff, error) {
	q := "INSERT INTO staffs (library_id, staff_id) VALUES ($1, $2) RETURNING library_id, user_id, created_at, updated_at, deleted_at;"
	args := []any{input.LibraryID, input.UserID}

	row := l.db.QueryRowContext(ctx, q, args...)

	var stf models.Staff
	err := row.Scan(&stf.ID, &stf.LibraryID, &stf.UserID, &stf.CreatedAt, &stf.UpdatedAt, &stf.DeletedAt)
	if err != nil {
		if strings.Contains(err.Error(), "staffs_library_id_fkey") {
			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "library not found", err)
		}
		if strings.Contains(err.Error(), "staffs_user_id_fkey") {
			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "user not found", err)
		}
		if strings.Contains(err.Error(), "staffs_pkey") {
			return nil, core.NewCoreError(core.ErrCodeDBDuplicate, "staff is already assigned", err)
		}
		return nil, core.NewCoreError(core.ErrCodeDBScan, "error creating staff", err)
	}

	return stf.ToCoreModel(), nil
}

func (l *LibraryAppDB) CountStaffs(ctx context.Context) (int, error) {
	var count int
	q := "SELECT count(id) from staff;"
	row := l.db.QueryRowContext(ctx, q)
	err := row.Scan(&count)
	return count, err
}
