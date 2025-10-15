package order

import "orderTracker/internal/service"

type Handler struct {
	Service service.OrderService
}

func NewOrderHandler(s service.OrderService) *Handler {
	return &Handler{Service: s}
}
