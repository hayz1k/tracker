package store

import "database/sql"

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(conn *sql.DB) *PostgresStore {
	return &PostgresStore{db: conn}
}
