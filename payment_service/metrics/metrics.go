package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	GRPCServerRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_server_requests_total",
			Help: "Total number of gRPC server requests",
		},
		[]string{"service", "method", "status"},
	)

	GRPCServerDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_server_duration_seconds",
			Help:    "Duration of gRPC server requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method"},
	)

	DBQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Duration of database queries",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation", "table"},
	)

	DBQueryErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_query_errors_total",
			Help: "Total number of database query errors",
		},
		[]string{"operation", "table"},
	)
)

func init() {
	prometheus.MustRegister(
		GRPCServerRequestsTotal,
		GRPCServerDuration,
		DBQueryDuration,
		DBQueryErrors,
	)
}
