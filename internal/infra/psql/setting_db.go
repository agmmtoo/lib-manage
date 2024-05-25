package psql

import (
	"context"
	"database/sql"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
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
