package postgresql

import (
	"database/sql"
	"fmt"
	"strings"
)

type DataContextGetter interface {
	// GetDataContext is used when we want to access underlying database for crud
	GetDataContext() *sql.DB
}

type PgController struct {
	db          *PgDatabase
	initialized bool
	config      Config
	testDB      bool
}

func NewPgController(dbConfig Config, testDB bool) *PgController {
	return &PgController{config: dbConfig, testDB: testDB}
}

// GetDataContext is used when we want to access underlying database for crud
func (d *PgController) GetDataContext() *sql.DB {

	return d.db.Database
}

const (
	DBCreateDMLIfNotExist         = "CREATE DATABASE %s "
	CreateUUIDExtensionIfNotExist = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	CreateSchemaIfNotExist        = `CREATE SCHEMA IF NOT EXISTS "%s";`
)

const (
	ErrorCodeDbExist          = "42P04"
	ErrorCodePermissionDenied = "42501"
)

func (d *PgController) Generate() error {
	sd := NewPgDatabase(d.config)
	d.db = sd
	err := sd.open()
	if err != nil {
		return err
	}
	tx, err := sd.GetDB().Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = sd.GetDB().Exec(fmt.Sprintf(DBCreateDMLIfNotExist, d.config.DBName))
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			tx.Rollback()
			return sd.Open()
		}
	}

	return nil
}

func (d *PgController) Init() error {
	if d.testDB {
		err := d.dropSchema()
		if err != nil {
			return err
		}
		err = d.createSchema()
		if err != nil {
			return err
		}
	}
	tx := NewTransaction(d.db.Database)
	err := tx.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return err
	}
	context := tx.GetDataContext()
	_, err = context.Exec(CreateUUIDExtensionIfNotExist)
	if err != nil {
		return err
	}
	if d.config.Schema != "" {
		_, err = context.Exec(fmt.Sprintf(CreateSchemaIfNotExist, d.config.Schema))
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err == nil {
		d.initialized = true
	}

	return err
}

func (d *PgController) dropSchema() error {
	tx := NewTransaction(d.db.Database)
	err := tx.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return err
	}
	context := tx.GetDataContext()
	context.Exec("DROP SCHEMA IF EXISTS public CASCADE;")
	return tx.Commit()
}

func (d *PgController) createSchema() error {
	tx := NewTransaction(d.db.Database)
	err := tx.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return err
	}
	context := tx.GetDataContext()
	context.Exec("CREATE SCHEMA IF NOT EXISTS public;")
	return tx.Commit()
}
