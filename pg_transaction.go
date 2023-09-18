package postgresql

import (
	"database/sql"
)

type Transaction struct {
	db       *sql.DB
	Tx       *sql.Tx
	commited bool
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{db: db}
}

// GetDataContext is used when we want to access underlying database for crud
func (t *Transaction) GetDataContext() *sql.DB {

	return t.db
}

func (t *Transaction) Begin() error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	t.Tx = tx
	return nil
}

func (t *Transaction) Rollback() error {
	if t.Tx == nil {
		return nil
	}
	err := t.Tx.Rollback()

	return err
}

func (t *Transaction) RollbackUnlessCommitted() {
	if !t.commited {
		_ = t.Rollback()
	}
}

func (t *Transaction) Commit() error {
	err := t.Tx.Commit()
	t.commited = true
	return err
}
