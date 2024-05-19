package psql

import (
	"context"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListUsers(ctx context.Context, input user.ListRequest) (*user.ListResponse, error) {
	q := "SELECT id, username, password, created_at, updated_at, deleted_at FROM \"user\";"
	args := []interface{}{}
	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*libraryapp.User
	for rows.Next() {
		var u libraryapp.User
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user.ListResponse{
		Users: users,
	}, nil
}

func (l *LibraryAppDB) GetUserByID(ctx context.Context, id int) (*libraryapp.User, error) {
	q := "SELECT id, username, password, created_at, updated_at, deleted_at FROM \"user\" WHERE id = $1;"
	args := []interface{}{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var u libraryapp.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (l *LibraryAppDB) CreateUser(ctx context.Context, input user.CreateRequest) (*libraryapp.User, error) {
	q := "INSERT INTO \"user\" (username, password) VALUES ($1, $2) RETURNING id, username, password, created_at, updated_at, deleted_at;"
	args := []interface{}{input.Username, input.Password}

	row := l.db.QueryRowContext(ctx, q, args...)

	var u libraryapp.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		if strings.Contains(err.Error(), "user_username_key") {
			return nil, libraryapp.CoreError{
				Code:   libraryapp.ECONFLICT,
				Reason: "username already exists",
			}
		}
		return nil, err
	}

	return &u, nil
}
