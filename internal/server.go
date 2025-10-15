package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"orderTracker/configs"
	"orderTracker/internal/http/handlers/order"
	"orderTracker/internal/http/handlers/site"
	"orderTracker/internal/http/handlers/updateorders"
	"orderTracker/internal/service"
	updateservice "orderTracker/internal/service/updateorders"
	"orderTracker/internal/store/postgres"
	"orderTracker/internal/woocommerce"
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

	wooClient := woocommerce.NewClient()

	updateOrdersSrv := updateservice.NewService(
		services.Site,
		services.Order,
		wooClient,
	)

	handlers := Handlers{
		Order:        order.NewOrderHandler(services.Order),
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
	r.Get("/orders/{id}", app.Handlers.Order.GetOrderByID)
	// r.Post("/orders{id}", app.Handlers.Order.CreateOrder)

	// Sites
	r.Get("/sites", app.Handlers.Site.GetSites)
	r.Post("/sites", app.Handlers.Site.PostSite)

	// Other
	r.Post("/update-orders", app.Handlers.UpdateOrders.UpdateOrders)

	return &Server{router: r}
}
