package psql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

// LibraryAppDB represents the database connection.
// Implements the core Storer interfaces.
type LibraryAppDB struct {
	db *sql.DB
}

// NewLibraryAppDB creates a new LibraryAppDB.
func NewLibraryAppDB(dataSourceName string) (*LibraryAppDB, error) {

	config, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		fmt.Println("New connection created")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDBFromPool(pool)

	// db, err := sql.Open("pgx", dataSourceName)
	// if err != nil {
	// 	return nil, err
	// }

	// Check if the database connection is working.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Migration
	libraryAppDB := &LibraryAppDB{db}
	err = libraryAppDB.Migrate()
	if err != nil {
		return nil, err
	}

	// Return the LibraryAppDB.
	return libraryAppDB, nil
}

// Close closes the database connection.
func (l *LibraryAppDB) Close() error {
	return l.db.Close()
}

func (l *LibraryAppDB) Ping(ctx context.Context) error {
	return l.db.PingContext(ctx)
}

// Migrate creates the tables in the database.
func (l *LibraryAppDB) Migrate() error {

	// NOTE: dependency for table creations
	fq, err := getSQLFileContent("create_function_update_updated_at_column.sql")
	if err != nil {
		return err
	}
	_, err = l.db.Exec(fq)
	if err != nil {
		return err
	}

	migrationFiles := []string{
		"create_table_book.sql",
		"create_table_user.sql",
		"create_table_library.sql",
		"create_table_staff.sql",
		"create_table_library_book.sql",
		"create_table_library_staff.sql",
		"create_table_loan.sql",
		"create_table_setting.sql",
	}

	var tq string

	for _, file := range migrationFiles {
		s, err := getSQLFileContent(file)
		if err != nil {
			return err
		}
		tq += s
	}

	// execute the migration SQL
	_, err = l.db.Exec(tq)
	return err
}

func getSQLFileContent(file string) (string, error) {
	migrationPath := "internal/infra/psql/migrations"
	b, err := os.ReadFile(path.Join(migrationPath, file))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
