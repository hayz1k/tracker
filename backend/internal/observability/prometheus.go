package observability

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	OrdersTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "orders_total",
			Help: "Total number of orders processed",
		},
	)
)

func Init() {
	prometheus.MustRegister(OrdersTotal)
}
