package client

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/metrics"
)

const (
	userService = "user_service"
	authService = "auth_service"
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
