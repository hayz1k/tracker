package prometheus

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (m *Metrics) MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			duration := time.Since(start).Seconds()
			path := normalizePath(r)
			method := r.Method

			statusClass := statusToClass(ww.Status())

			m.HTTPRequestsTotal.WithLabelValues(method, path, statusClass).Inc()
			m.HTTPRequestDuration.WithLabelValues(method, path).Observe(duration)
		}()

		next.ServeHTTP(ww, r)
	})
}

func statusToClass(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "2xx"
	case code >= 300 && code < 400:
		return "3xx"
	case code >= 400 && code < 500:
		return "4xx"
	case code >= 500:
		return "5xx"
	default:
		return strconv.Itoa(code)
	}
}

func normalizePath(r *http.Request) string {

	routePattern := chi.RouteContext(r.Context()).RoutePattern()
	if routePattern != "" {
		return routePattern
	}

	// Если не chi, fallback на ручные замены
	path := r.URL.Path

	// Пример для /orders/{id}
	if strings.HasPrefix(path, "/api/orders/") {
		return "/api/orders/:id"
	}
	if strings.HasPrefix(path, "/api/sites/") {
		return "/api/sites/:id"
	}

	return path
}
