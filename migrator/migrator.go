package migrator

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	db         *sql.DB
	migrations *migrate.FileMigrationSource
}

// TODO - set migration table name
// TODO - add limit to Up and Down method

func New(db *sql.DB, fileMigrationSource string) Migrator {
	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: fileMigrationSource,
	}

	return Migrator{db: db, dialect: "postgres", migrations: migrations}
}

func (m Migrator) Up() {

	n, err := migrate.Exec(m.db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	n, err := migrate.Exec(m.db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %v", err))
	}
	fmt.Printf("Rollbacked %d migrations!\n", n)
}

func (m Migrator) Status() {
	// TODO - add status
}
