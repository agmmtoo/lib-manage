package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/infra/psql/migrations"
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
	logger.Info("Migrating the database")
	_, err := l.db.Exec(migrations.CreateFunctionUpdateUpdatedAtColumn)
	if err != nil {
		return err
	}

	qs := []string{
		migrations.CreateTableBook,
		migrations.CreateTableUser,
		migrations.CreateTableLibrary,
		migrations.CreateTableStaff,
		migrations.CreateTableLibraryBook,
		migrations.CreateTableLibraryStaff,
		migrations.CreateTableLoan,
		migrations.CreateTableSetting,
	}

	// execute the migration SQL
	_, err = l.db.Exec(strings.Join(qs, "\n"))
	return err
}
