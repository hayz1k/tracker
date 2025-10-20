package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

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
	r.Use(app.Metrics.MetricsMiddleware)

	RegisterRoutes(r, &app.Handlers)

	return &Server{router: r}
}
