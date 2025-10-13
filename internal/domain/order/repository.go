package order

import "context"

type Store interface {
	GetByTrackNumber(ctx context.Context, trackNumber string) (*Order, error)
	GetByID(ctx context.Context, id int) (*Order, error)
	Save(ctx context.Context, order *Order) error
	Exists(ctx context.Context, trackNumber string) (bool, error)
	// updateStatus - подумать
	// UpdateStatus(ctx context.Context, status domain.Order) error
}
