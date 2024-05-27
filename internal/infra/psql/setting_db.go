package psql

import (
	"context"
	"database/sql"
	"fmt"

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

	qb := &QueryBuilder{
		Table:        "setting",
		ParamCounter: 1,
	}
	if len(input.IDs) > 0 {
		qb.AddClause("id = ANY($%d)", pq.Array(input.IDs))
	}
	if len(input.Key) > 0 {
		qb.AddClause("key ILIKE $%d", fmt.Sprintf("%%%s%%", input.Key))
	}
	if len(input.LibraryIDs) > 0 {
		qb.AddClause("library_id = ANY($%d)", pq.Array(input.LibraryIDs))
	}
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)

	query, params := qb.Build()

	rows, err := l.db.QueryContext(ctx, query, params...)
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
