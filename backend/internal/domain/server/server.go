package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"orderTracker/configs"
	"orderTracker/internal/http/handlers/login"
	"orderTracker/internal/http/handlers/order"
	"orderTracker/internal/http/handlers/site"
	"orderTracker/internal/http/handlers/updateorders"
	"orderTracker/internal/service"
	updateservice "orderTracker/internal/service/updateorders"
	"orderTracker/internal/store/postgres"
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
}

func NewApp(cfg *configs.Config, store *postgres.Store) *App {
	services := Services{
		Order: service.NewOrderService(store.Orders()),
		Site:  service.NewSiteService(store.Sites()),
	}

	wooClient := updateservice.NewClient()

	updateOrdersSrv := updateservice.NewService(
		services.Site,
		services.Order,
		wooClient,
	)

	handlers := Handlers{
		Order:        order.NewOrderHandler(services.Order, services.Site),
		Site:         site.NewSiteHandler(services.Site),
		UpdateOrders: updateorders.NewHandler(updateOrdersSrv),
	}

	return &App{
		Handlers: handlers,
		Services: services,
		Cfg:      cfg,
	}
}

type Server struct {
	router *chi.Mux
}

func (s *Server) Run(address string) error {
	return http.ListenAndServe(address, s.router)
}

func NewServer(app *App) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Orders
	r.Get("/api/orders/{id}", app.Handlers.Order.GetOrderByID)
	r.Get("/api/orders", app.Handlers.Order.GetOrders)
	r.Post("/api/orders", app.Handlers.Order.PostOrder)
	// r.Delete("/api/orders/{id}", app.Handlers.Order.DeleteOrder)
	// r.Post("/orders{id}", app.Handlers.Order.CreateOrder)

	// Sites
	r.Get("/api/sites", app.Handlers.Site.GetSites)
	r.Post("/api/sites", app.Handlers.Site.PostSite)
	// r.Put("/api/sites", app.Handlers.Site.PutSite)
	// r.Delete("/api/sites/{id}", app.Handlers.Site.DeleteSite)

	// Other
	r.Post("/update-orders", app.Handlers.UpdateOrders.UpdateOrders)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	r.Post("/api/auth/login", login.ServeHTTP)
	return &Server{router: r}
}
