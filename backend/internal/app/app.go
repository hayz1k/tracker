package app

import (
	"orderTracker/configs"
	"orderTracker/internal/adapter/delivery/http/handlers/order"
	"orderTracker/internal/adapter/delivery/http/handlers/site"
	"orderTracker/internal/adapter/delivery/http/handlers/updateorders"
	updateservice "orderTracker/internal/infrastructure/httpclient/woocommerce"
	"orderTracker/internal/infrastructure/observability/prometheus"
	"orderTracker/internal/infrastructure/store/postgres"
	"orderTracker/internal/service"
	updateorders2 "orderTracker/internal/usecases/updateorders"
)

type Handlers struct {
	Order        *order.Handler
	Site         *site.Handler
	UpdateOrders *updateorders.Handler
}

type Services struct {
	Order service.OrderService
	Site  service.SiteService
}

type App struct {
	Handlers Handlers
	Services Services
	Cfg      *configs.Config
	Metrics  *prometheus.Metrics
}

func NewApp(cfg *configs.Config, store *postgres.Store) *App {
	services := Services{
		Order: service.NewOrderService(store.Orders()),
		Site:  service.NewSiteService(store.Sites()),
	}

	metrics := prometheus.NewMetrics()

	wooClient := updateservice.NewClient()

	updateOrdersSrv := updateorders2.NewService(
		services.Site,
		services.Order,
		wooClient,
	)

	handlers := Handlers{
		Order:        order.NewOrderHandler(services.Order, services.Site, metrics),
		Site:         site.NewSiteHandler(services.Site),
		UpdateOrders: updateorders.NewHandler(updateOrdersSrv),
	}

	return &App{
		Handlers: handlers,
		Services: services,
		Cfg:      cfg,
		Metrics:  metrics,
	}
}
