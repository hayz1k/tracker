package postgres

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"orderTracker/configs"
	"orderTracker/internal/infrastructure/store/postgres/orderstore"
	"orderTracker/internal/infrastructure/store/postgres/sitestore"
	"time"
)

type Store struct {
	db *sql.DB
}

func NewPostgresStore(cfg *configs.Config) (*Store, error) {
	db, err := sql.Open("pgx", cfg.DSN)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(5 * time.Minute)
	
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
