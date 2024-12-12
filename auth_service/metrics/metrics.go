package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// GRPCServerRequestsTotal is a counter vector for total grpc requests number
	GRPCServerRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_server_requests_total",
			Help: "Total number of gRPC server requests",
		},
		[]string{"service", "method", "status"},
	)

	// GRPCServerDuration is a histogram vector for grpc requests duration
	GRPCServerDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_server_duration_seconds",
			Help:    "Duration of gRPC server requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method"},
	)

	// RedisQueryDuration is a histrogram vector for redis requests duration
	RedisQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "redis_query_duration_seconds",
			Help:    "Duration of Redis queries",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)

	// RedisQueryErrors is a counter vector for redis requests errors
	RedisQueryErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "redis_query_errors_total",
			Help: "Total number of redis query errors",
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(
		GRPCServerRequestsTotal,
		GRPCServerDuration,
		RedisQueryDuration,
		RedisQueryErrors,
	)
}
