package order

import (
	"orderTracker/internal/infrastructure/observability/prometheus"
	"orderTracker/internal/service"
)

type Handler struct {
	Service     service.OrderService
	SiteService service.SiteService
	Metrics     *prometheus.Metrics
}

func NewOrderHandler(s service.OrderService, ss service.SiteService, m *prometheus.Metrics) *Handler {
	return &Handler{Service: s, SiteService: ss, Metrics: m}
}
