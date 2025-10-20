package prometheus

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HTTPRequestsTotal   *prometheus.CounterVec
	HTTPRequestDuration *prometheus.HistogramVec
	DBStats             *prometheus.GaugeVec
	BusinessOrders      *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	m := &Metrics{
		HTTPRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests grouped by method, path, and status class",
			},
			[]string{"method", "path", "status_class"},
		),
		HTTPRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "Histogram of request durations",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "path"},
		),
		DBStats: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "db_pool_connections",
				Help: "Database connection pool statistics",
			},
			[]string{"state"},
		),
		BusinessOrders: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "orders_total",
				Help: "Total number of orders created/failed by status",
			},
			[]string{"status"},
		),
	}

	prometheus.MustRegister(
		m.HTTPRequestsTotal,
		m.HTTPRequestDuration,
		m.DBStats,
		m.BusinessOrders,
	)

	return m
}
