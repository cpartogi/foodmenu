package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cpartogi/foodmenu/module/menu"
)

// SQLStore provides all functions to execute db queries and transactions.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) menu.Repository {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTX executes a function within a database transaction
func (s *SQLStore) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
