package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HTTPRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests with method, path and status",
		},
		[]string{"method", "path", "status"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	GRPCClientRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_client_requests_total",
			Help: "Total number of gRPC client requests",
		},
		[]string{"service", "method", "status"},
	)

	GRPCClientDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_client_duration_seconds",
			Help:    "Duration of gRPC client requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method"},
	)
)

func init() {
	prometheus.MustRegister(
		HTTPRequestTotal,
		HTTPRequestDuration,
		GRPCClientRequestsTotal,
		GRPCClientDuration,
	)
}
