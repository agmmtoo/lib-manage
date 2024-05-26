package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/core/setting"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
	"github.com/lib/pq"
)

func (l *LibraryAppDB) GetSettingValue(ctx context.Context, libraryID int, key string) (string, error) {
	q := "SELECT value FROM setting WHERE library_id = $1 AND key = $2;"
	args := []any{libraryID, key}

	row := l.db.QueryRowContext(ctx, q, args...)

	var value string
	err := row.Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", libraryapp.NewCoreError(libraryapp.ErrCodeDBNotFound, "setting not found", err)
		}
		return "", libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error getting setting", err)
	}

	return value, nil
}

func (l *LibraryAppDB) ListSettings(ctx context.Context, input setting.ListRequest) (*setting.ListResponse, error) {

	finalQuery, params := _buildListQuery(input)

	fmt.Println(finalQuery, params)

	// q := "SELECT id, library_id, key, value, created_at, updated_at, deleted_at FROM setting;"
	rows, err := l.db.QueryContext(ctx, finalQuery, params...)
	if err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing settings", err)
	}
	defer rows.Close()

	var settings []*libraryapp.Setting
	for rows.Next() {
		var s libraryapp.Setting
		err := rows.Scan(&s.ID, &s.LibraryID, &s.Key, &s.Value, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBScan, "error scanning setting", err)
		}
		settings = append(settings, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error listing settings", err)
	}

	return &setting.ListResponse{
		Settings: settings,
	}, nil

}

func (l *LibraryAppDB) UpdateSettingValues(ctx context.Context, input []setting.UpdateRequest) ([]*libraryapp.Setting, error) {
	q := "UPDATE setting SET value = $1 WHERE id = $2 RETURNING id, library_id, key, value, created_at, updated_at, deleted_at;"
	var result []*libraryapp.Setting
	for _, st := range input {
		var s libraryapp.Setting
		args := []any{st.Value, st.ID}
		row := l.db.QueryRowContext(ctx, q, args...)

		err := row.Scan(&s.ID, &s.LibraryID, &s.Key, &s.Value, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			return nil, libraryapp.NewCoreError(libraryapp.ErrCodeDBQuery, "error updating setting", err)
		}
		result = append(result, &s)
	}

	return result, nil
}

func _buildListQuery(req setting.ListRequest) (string, []interface{}) {
	var (
		clauses      []string
		params       []any
		paramCounter = 1
	)

	if len(req.IDs) > 0 {
		clauses = append(clauses, fmt.Sprintf("id = ANY($%d)", paramCounter))
		params = append(params, pq.Array(req.IDs))
		paramCounter++
	}
	if len(req.Key) > 0 {
		clauses = append(clauses, fmt.Sprintf("key LIKE $%d", paramCounter))
		params = append(params, fmt.Sprintf("%%%s%%", req.Key))
		paramCounter++
	}
	if len(req.LibraryIDs) > 0 {
		clauses = append(clauses, fmt.Sprintf("library_id = ANY($%d)", paramCounter))
		params = append(params, pq.Array(req.LibraryIDs))
		paramCounter++
	}

	query := "SELECT * FROM setting"

	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	if req.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", len(params)+1)
		params = append(params, req.Limit)
	}
	if req.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", len(params)+1)
		params = append(params, req.Offset)
	}

	return query, params
}
