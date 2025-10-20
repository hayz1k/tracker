package orderstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/order"
	"strings"
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

	query := "SELECT * FROM orders WHERE order_id = $1"
	err := o.db.QueryRow(query, orderID).
		Scan(&result.OrderID, &result.FirstName, &result.SecondName,
			&result.DeliveryAddress, &result.Total, &result.CurrentStatus,
			&result.Created, &result.TrackNumber, &result.SiteID)
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

func (o *OrderStore) ListOrders(ctx context.Context, page, limit int, f *order.OrderFilter) ([]*order.Order, error) {
	offset := (page - 1) * limit

	baseQuery := `
		SELECT 
			order_id, first_name, second_name, 
			address, total, status, 
			created_at, track_number, site_id
		FROM orders
	`

	var whereClauses []string
	var args []interface{}
	argIndex := 1

	if f != nil {
		if f.Status != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("current_status = $%d", argIndex))
			args = append(args, f.Status)
			argIndex++
		}
		if f.SiteID != 0 {
			whereClauses = append(whereClauses, fmt.Sprintf("site_id = $%d", argIndex))
			args = append(args, f.SiteID)
			argIndex++
		}
		if f.Search != "" {
			whereClauses = append(whereClauses, fmt.Sprintf(
				"(first_name ILIKE $%d OR second_name ILIKE $%d OR track_number ILIKE $%d)",
				argIndex, argIndex, argIndex,
			))
			args = append(args, "%"+f.Search+"%")
			argIndex++
		}
	}

	// добавляем WHERE, если есть фильтры
	if len(whereClauses) > 0 {
		baseQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	baseQuery += fmt.Sprintf(" ORDER BY order_id DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	rows, err := o.db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*order.Order
	for rows.Next() {
		var result order.Order
		if err := rows.Scan(
			&result.OrderID,
			&result.FirstName,
			&result.SecondName,
			&result.DeliveryAddress,
			&result.Total,
			&result.CurrentStatus,
			&result.Created,
			&result.TrackNumber,
			&result.SiteID,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &result)
	}
	return orders, nil
}

func (o *OrderStore) Count(ctx context.Context) (int, error) {
	var total int

	query := "SELECT COUNT(*) FROM orders"
	err := o.db.QueryRow(query).Scan(&total)
	if err != nil {
		log.Info().Err(err).Msg("error getting orders count")
		if errors.Is(err, sql.ErrNoRows) {
			return -1, errors.New("orders not found")
		}
		return -1, errors.New("internal error")
	}
	log.Info().Msg("success getting orders count")
	return total, nil
}

func (o *OrderStore) GetByTrackNumber(ctx context.Context, trackNumber string) (*order.Order, error) {
	log.Info().Msg("starting getting order by trackNumber")
	var result order.Order

	query := "SELECT * FROM orders WHERE track_number = $1"
	err := o.db.QueryRow(query, trackNumber).
		Scan(&result.OrderID, &result.FirstName, &result.SecondName,
			&result.DeliveryAddress, &result.Total, &result.CurrentStatus,
			&result.Created, &result.TrackNumber, &result.SiteID)
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
		"(order_id, first_name, second_name, address, total, status, created_at, track_number, site_id) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	_, err := o.db.Exec(query, order.OrderID, order.FirstName, order.SecondName,
		order.DeliveryAddress, order.Total, order.CurrentStatus, order.Created, order.TrackNumber, order.SiteID)
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

func (o *OrderStore) Exists(ctx context.Context, orderID int) (bool, error) {
	var tmp int
	query := "SELECT 1 FROM orders WHERE order_id = $1 LIMIT 1"

	err := o.db.QueryRowContext(ctx, query, orderID).Scan(&tmp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		log.Error().Err(err).Msgf("failed to check existence of order %d", orderID)
		return false, fmt.Errorf("internal error: %w", err)
	}

	return true, nil
}

func (o *OrderStore) Delete(ctx context.Context, id int) (bool, error) {
	query := "DELETE FROM orders WHERE order_id = $1"
	res, err := o.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to delete order %d", id)
		return false, fmt.Errorf("internal error: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msgf("failed to get rows affected for order %d", id)
		return false, fmt.Errorf("internal error: %w", err)
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
