package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/order"
)

type OrderService interface {
	GetByID(ctx context.Context, id int) (*order.Order, error)
	Save(ctx context.Context, order *order.Order) error
	ListOrders(ctx context.Context, page, limit int, f *order.OrderFilter) ([]*order.Order, error)
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, orderID int) (bool, error)
}

type orderService struct {
	store order.Store
}

func NewOrderService(store order.Store) OrderService {
	return &orderService{
		store: store,
	}
}

func (os *orderService) GetByID(ctx context.Context, id int) (*order.Order, error) {
	return os.store.GetByID(ctx, id)
}

func (os *orderService) Save(ctx context.Context, order *order.Order) error {
	log.Info().Msg("start save order service")
	if err := order.Validate(); err != nil {
		log.Error().Msg("invalid order")
		return err
	}
	return os.store.Save(ctx, order)
}

func (os *orderService) Exists(ctx context.Context, orderID int) (bool, error) {
	return os.store.Exists(ctx, orderID)
}

func (os *orderService) ListOrders(ctx context.Context, page, limit int, f *order.OrderFilter) ([]*order.Order, error) {
	return os.store.ListOrders(ctx, page, limit, f)
}

func (os *orderService) Count(ctx context.Context) (int, error) {
	return os.store.Count(ctx)
}
