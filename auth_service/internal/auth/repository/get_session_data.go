package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/metrics"
	"github.com/rs/zerolog/log"
)

// GetSessionData gets session info from Redis
func (ar *AuthRepository) GetSessionData(ctx context.Context, cookie string) (string, error) {
	var userID string
	start := time.Now()
	logger := log.Ctx(ctx)

	err := ar.Redis.Get(ctx, cookie).Scan(&userID)

	duration := time.Since(start).Seconds()
	metrics.RedisQueryDuration.WithLabelValues("get_from_redis").Observe(duration)

	if err != nil {
		metrics.RedisQueryErrors.WithLabelValues("get_from_redis").Inc()
		errMsg := fmt.Errorf("redis: cannot get cookie from redis - %w", err)
		logger.Error().Err(errMsg).Msg("redis_get_error")

		return "", errMsg
	}

	logger.Info().Msg("redis: successfully get info from cookie")
	return userID, nil
}
