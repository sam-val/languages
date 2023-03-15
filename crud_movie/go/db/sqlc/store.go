package db

import "database/sql"

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(conn *sql.DB) *Store {
	return &Store{
		Queries: New(conn),
		db: conn,
	}
}
