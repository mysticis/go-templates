package db

import "database/sql"

type Store interface {
	Querier
}

//Store provides all functions to execute DB transactions and queries
type SQLStore struct {
	*Queries
	db *sql.DB
}

//NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
