package postgres

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"orderTracker/configs"
	"orderTracker/internal/store/postgres/orderstore"
	"orderTracker/internal/store/postgres/sitestore"
)

type Store struct {
	db *sql.DB
}

func NewPostgresStore(cfg *configs.Config) (*Store, error) {
	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Orders() *orderstore.OrderStore {
	return orderstore.NewOrderStore(s.db)
}

func (s *Store) Sites() *sitestore.SiteStore {
	return sitestore.NewSiteStore(s.db)
}

func (s *Store) Close() error {
	return s.db.Close()
}
