package psql

import (
	"context"

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

	return nil, nil
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
