package psql

import (
	"context"
	"database/sql"
	"fmt"

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
	// _, err := l.db.Exec(migrations.CreateFunctionUpdateUpdatedAtColumn)
	// if err != nil {
	// 	return err
	// }

	qs := []string{
		migrations.CreateTableUsers,
		migrations.CreateTableLibraries,
		migrations.CreateTableCategories,
		migrations.CreateTableSubCategories,
		migrations.CreateTableBooks,
		migrations.CreateTableLibrariesBooks,
		migrations.CreateTableStaffs,
		migrations.CreateTableMemberships,
		migrations.CreateTableUsersMemberships,
		migrations.CreateTableLoans,
		migrations.CreateTableSetting,
	}
	// execute the migration SQL
	tx, err := l.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	for _, q := range qs {
		if _, err := tx.ExecContext(context.Background(), q); err != nil {
			logger.Error("error executing migration", "error", err, "query", q)
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		logger.Error("error committing migration", "error", err)
		return err
	}
	logger.Info("Migration successful")
	return nil
}
