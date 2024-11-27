package client

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/metrics"
)

const (
	userClient  = "user_client"
	authClient  = "auth_client"
	movieClient = "movie_client"
)

func status(err error) string {
	if err != nil {
		return "error"
	}

	return "success"
}

func saveMetric(start time.Time, service string, method string, err error) {
	duration := time.Since(start).Seconds()
	metrics.GRPCClientRequestsTotal.WithLabelValues(service, method, status(err)).Inc()
	metrics.GRPCClientDuration.WithLabelValues(service, method).Observe(duration)
}
