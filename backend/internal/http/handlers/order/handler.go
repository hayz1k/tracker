package order

import "orderTracker/internal/service"

type Handler struct {
	Service     service.OrderService
	SiteService service.SiteService
}

func NewOrderHandler(s service.OrderService, ss service.SiteService) *Handler {
	return &Handler{Service: s, SiteService: ss}
}
