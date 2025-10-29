package order

import (
	"context"
)

type Store interface {
	Save(ctx context.Context, order *Order) error
	GetByID(ctx context.Context, id int) (*Order, error)
	GetByTrackNumber(ctx context.Context, trackNumber string) (*Order, error)
	// UpdateStatus(ctx context.Context, id int, status string) error

	ListOrders(ctx context.Context, page, limit int, f *OrderFilter) ([]*Order, error)
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, orderID int, siteID int) (bool, error)
	Delete(ctx context.Context, id int) (bool, error)
	Update(ctx context.Context, ord *Order) error
}
