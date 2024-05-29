package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListUsers(ctx context.Context, input user.ListRequest) (*user.ListResponse, error) {
	qb := &QueryBuilder{
		Table:        "\"user\"",
		ParamCounter: 1,
		Cols:         []string{"id", "username", "created_at", "updated_at", "deleted_at"},
	}
	if len(input.IDs) > 0 {
		qb.AddClause("id = ANY($%d)", input.IDs)
	}
	if len(input.Username) > 0 {
		qb.AddClause("username ILIKE $%d", fmt.Sprintf("%%%s%%", input.Username))
	}
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)
	q, args := qb.Build()

	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing users", err)
	}

	defer rows.Close()

	var users []*libraryapp.User
	for rows.Next() {
		var u libraryapp.User
		err := rows.Scan(&u.ID, &u.Username, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error scanning user", err)
		}
		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing users", err)
	}

	return &user.ListResponse{
		Users: users,
	}, nil
}

func (l *LibraryAppDB) GetUserByID(ctx context.Context, id int) (*libraryapp.User, error) {
	q := "SELECT id, username, password, created_at, updated_at, deleted_at FROM \"user\" WHERE id = $1;"
	args := []any{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var u libraryapp.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "user not found", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error getting user", err)
	}

	return &u, nil
}

func (l *LibraryAppDB) CreateUser(ctx context.Context, input user.CreateRequest) (*libraryapp.User, error) {
	q := "INSERT INTO \"user\" (username, password) VALUES ($1, $2) RETURNING id, username, password, created_at, updated_at, deleted_at;"
	args := []any{input.Username, input.Password}

	row := l.db.QueryRowContext(ctx, q, args...)

	var u libraryapp.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		if strings.Contains(err.Error(), "user_username_key") {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBDuplicate, "username already exists", err)
		}
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error creating user", err)
	}

	return &u, nil
}
