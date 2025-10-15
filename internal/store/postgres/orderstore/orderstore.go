package orderstore

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/order"
)

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{db: db}
}

func (o *OrderStore) GetByID(ctx context.Context, orderID int) (*order.Order, error) {
	log.Info().Msg("starting getting order by id")
	var result order.Order
	var dbID int

	query := "SELECT * FROM orders WHERE order_id = $1"
	err := o.db.QueryRow(query, orderID).
		Scan(&dbID, &result.OrderID, &result.FirstName, &result.SecondName,
			&result.DeliveryAddress, &result.TrackNumber, &result.CurrentStatus,
			&result.Total, &result.Created, &result.SiteID)
	if err != nil {
		log.Error().Err(err).Msg("error getting result by id")
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, errors.New("internal error")
	}
	log.Info().Msg("success getting order by id")
	return &result, nil
}

func (o *OrderStore) GetByTrackNumber(ctx context.Context, trackNumber string) (*order.Order, error) {
	log.Info().Msg("starting getting order by trackNumber")
	var result order.Order
	var dbID int

	query := "SELECT * FROM orders WHERE track_number = $1"
	err := o.db.QueryRow(query, trackNumber).
		Scan(&dbID, &result.OrderID, &result.FirstName, &result.SecondName,
			&result.DeliveryAddress, &result.TrackNumber, &result.CurrentStatus,
			&result.Total, &result.Created, &result.SiteID)
	if err != nil {
		log.Info().Err(err).Msg("error getting order by track number")
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, errors.New("internal error")
	}
	log.Info().Msg("success getting order by track number")
	return &result, nil
}

func (o *OrderStore) Save(ctx context.Context, order *order.Order) error {
	log.Info().Msg("saving order")
	query := "INSERT INTO orders " +
		"(order_id, first_name, last_name, address, status, track_number, total, created_at, site_id) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	_, err := o.db.Exec(query, order.OrderID, order.FirstName, order.SecondName,
		order.DeliveryAddress, order.CurrentStatus, order.TrackNumber, order.Total, order.Created, order.SiteID)
	if err != nil {
		log.Error().Err(err).Msg("error adding order")
		return errors.New("internal error")
	}
	log.Info().Msg("successfully added order")
	return nil
}

func (o *OrderStore) UpdateStatus(ctx context.Context, id int, status string) error {
	return nil
}

func (o *OrderStore) StatusesByID(ctx context.Context, id int) {
	panic("implement me")
}

func (o *OrderStore) Exists(ctx context.Context, trackNumber string) (bool, error) {
	return false, nil
}
