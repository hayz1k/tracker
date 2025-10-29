package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/status"
)

type StatusStore struct {
	db *sql.DB
}

func NewStatusStore(db *sql.DB) *StatusStore {
	return &StatusStore{db: db}
}

func (s *StatusStore) Save(ctx context.Context, status *status.Status) error {
	query := "INSERT INTO statuses(order_id, status, is_custom, created_at)" +
		"VALUES ($1, $2, $3, $4)"
	_, err := s.db.ExecContext(ctx, query, status.OrderID, status.Status, status.IsCustom, status.CreatedAt)
	if err != nil {
		log.Error().Err(err).Msg("error adding status")
		return errors.New("internal error")
	}
	log.Info().Msg("successfully added status")
	return nil
}

func (s *StatusStore) GetByID(ctx context.Context, id int) ([]*status.Status, error) {
	var result []*status.Status

	query := "SELECT id, order_id, site_id, status, is_custom, created_at FROM statuses WHERE id = $1"
	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var st status.Status
		if err := rows.Scan(&st.ID, &st.OrderID, &st.SiteID, &st.Status, &st.IsCustom, &st.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &st)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *StatusStore) GetByData(ctx context.Context, orderID, siteID int) ([]*status.Status, error) {
	var result []*status.Status

	query := "SELECT id, order_id, site_id, status, is_custom, created_at FROM statuses WHERE order_id = $1 AND site_id = $2"
	rows, err := s.db.QueryContext(ctx, query, orderID, siteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var st status.Status
		if err := rows.Scan(&st.ID, &st.OrderID, &st.SiteID, &st.Status, &st.IsCustom, &st.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &st)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
