package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"orderTracker/internal/adapter/delivery/http/handlers/login"
)

func RegisterRoutes(r *chi.Mux, handler *Handlers) {

	// Orders
	r.Get("/api/orders/count", handler.Order.GetOrdersCount)
	r.Get("/api/orders/{id}", handler.Order.GetOrderByID)
	r.Get("/api/orders", handler.Order.GetOrders)
	r.Post("/api/orders", handler.Order.PostOrder)
	r.Delete("/api/orders/{id}", handler.Order.DeleteOrder)
	// r.Post("/orders{id}", handler.Order.CreateOrder)

	// Sites
	r.Get("/api/sites", handler.Site.GetSites)
	r.Post("/api/sites", handler.Site.PostSite)
	// r.Put("/api/sites", handler.Site.PutSite)
	// r.Delete("/api/sites/{id}", handler.Site.DeleteSite)

	// Other
	r.Post("/update-orders", handler.UpdateOrders.UpdateOrders)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	r.Post("/api/auth/login", login.ServeHTTP)
}
