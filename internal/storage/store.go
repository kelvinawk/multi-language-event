package store

import "database/sql"

// store all repo
type Store struct {
}

func NewStore(db *sql.DB) *Store {
	return &Store{}
}
