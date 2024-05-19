package psql

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/core/library"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

func (l *LibraryAppDB) ListLibraries(ctx context.Context, input library.ListRequest) (*library.ListResponse, error) {
	q := "SELECT id, name, created_at, updated_at, deleted_at FROM library;"
	args := []any{}
	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var libraries []*libraryapp.Library
	for rows.Next() {
		var lib libraryapp.Library
		err := rows.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
		if err != nil {
			return nil, err
		}
		libraries = append(libraries, &lib)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &library.ListResponse{
		Libraries: libraries,
	}, nil
}

func (l *LibraryAppDB) GetLibraryByID(ctx context.Context, id int) (*libraryapp.Library, error) {
	q := "SELECT id, name, created_at, updated_at, deleted_at FROM library WHERE id = $1;"
	args := []interface{}{id}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lib libraryapp.Library
	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &lib, nil
}

func (l *LibraryAppDB) CreateLibrary(ctx context.Context, input library.CreateRequest) (*libraryapp.Library, error) {
	q := "INSERT INTO library (name) VALUES ($1) RETURNING id, name, created_at, updated_at, deleted_at;"
	args := []any{input.Name}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lib libraryapp.Library
	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedAt, &lib.UpdatedAt, &lib.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &lib, nil
}
