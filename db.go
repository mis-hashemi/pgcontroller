package postgresql

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type PgDatabase struct {
	DBConfig Config
	Database *sql.DB
}

func (m *PgDatabase) GetDB() *sql.DB {
	return m.Database
}

func (md *PgDatabase) Open() error {
	return md.open()
}

func NewPgDatabase(config Config) *PgDatabase {
	md := PgDatabase{}
	md.DBConfig = config
	return &md
}

func (md *PgDatabase) open() error {
	db, err := sql.Open("postgres", md.DBConfig.DSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	md.Database = db
	return nil
}
