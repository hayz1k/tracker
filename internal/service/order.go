package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/order"
)

type OrderService interface {
	SaveOrder(ctx context.Context, order *order.Order) error
	GetOrderByID(ctx context.Context, id int) (*order.Order, error)
}

type orderService struct {
	store order.Store
}

func NewOrderService(store order.Store) OrderService {
	return &orderService{
		store: store,
	}
}

func (os *orderService) GetOrderByID(ctx context.Context, id int) (*order.Order, error) {
	return os.store.GetByID(ctx, id)
}

func (os *orderService) SaveOrder(ctx context.Context, order *order.Order) error {
	log.Info().Msg("start save order service")
	if err := order.Validate(); err != nil {
		log.Error().Msg("invalid order")
		return err
	}
	return os.store.Save(ctx, order)
}
