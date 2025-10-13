package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/order"
	errors2 "orderTracker/internal/errors"
)

type OrderStore struct {
	db *sql.DB
}

func (o *OrderStore) GetByID(ctx context.Context, orderID int) (*order.Order, error) {
	log.Info().Msg("starting getting result by id")
	var result order.Order
	var dbID int

	query := "SELECT * FROM users WHERE order_id = $1"
	err := o.db.QueryRow(query, orderID).
		Scan(&dbID, &result.OrderID, &result.FirstName, &result.SecondName,
			&result.DeliveryAddress, &result.CurrentStatus, &result.Created,
			&result.TrackNumber)
	if err != nil {
		log.Info().Err(err).Msg("error getting result by id")
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.ErrOrderNotFound
		}
		return nil, errors2.ErrInternal
	}
	log.Info().Msg("success getting result by id")
	return &result, nil
}

func (o *OrderStore) GetByTrackNumber(ctx context.Context, trackNumber string) (*order.Order, error) {
	log.Info().Msg("starting getting order by trackNumber")
	var order order.Order
	var dbID int

	query := "SELECT * FROM orders WHERE track_number = $1"
	err := o.db.QueryRow(query, trackNumber).
		Scan(&dbID, &order.OrderID, &order.FirstName, &order.SecondName,
			&order.DeliveryAddress, &order.CurrentStatus, &order.Created,
			&order.TrackNumber)
	if err != nil {
		log.Info().Err(err).Msg("error getting order by track number")
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.ErrOrderNotFound
		}
		return nil, errors2.ErrInternal
	}
	log.Info().Msg("success getting order by track number")
	return &order, nil
}

func (o *OrderStore) Save(ctx context.Context, order *order.Order) error {
	log.Info().Msg("saving order")
	query := "INSERT INTO orders " +
		"(first_name, second_name, address, status, created, track_number) " +
		"VALUES($1,$2,$3,$4,$5,$6)"
	_, err := o.db.Exec(query, order.OrderID, order.FirstName, order.SecondName,
		order.DeliveryAddress, order.CurrentStatus, order.Created, order.TrackNumber)
	if err != nil {
		log.Error().Err(err).Msg("error adding order")
		return errors2.ErrInternal
	}
	log.Info().Msg("successfully added order")
	return nil
}

func (o *OrderStore) Exists(ctx context.Context, trackNumber string) (bool, error) {
	return false, nil
}
