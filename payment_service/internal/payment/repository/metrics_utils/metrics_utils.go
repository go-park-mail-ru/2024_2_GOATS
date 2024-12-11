package metricsutils

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/metrics"
)

// SaveSuccessMetric creates new DB success metric
func SaveSuccessMetric(start time.Time, operation string, table string) {
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues(operation, table).Observe(duration)
}

// SaveErrorMetric creates new DB error metric
func SaveErrorMetric(operation string, table string) {
	metrics.DBQueryErrors.WithLabelValues(operation, table).Inc()
}
