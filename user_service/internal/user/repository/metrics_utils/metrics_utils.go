package metricsutils

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/metrics"
)

func SaveSuccessMetric(start time.Time, operation string, table string) {
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues(operation, table).Observe(duration)
}

func SaveErrorMetric(start time.Time, operation string, table string) {
	metrics.DBQueryErrors.WithLabelValues(operation, table).Inc()
}
