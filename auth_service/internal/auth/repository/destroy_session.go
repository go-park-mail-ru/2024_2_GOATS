package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/metrics"
	"github.com/rs/zerolog/log"
)

func (ar *AuthRepository) DestroySession(ctx context.Context, cookie string) error {
	logger := log.Ctx(ctx)
	start := time.Now()

	_, err := ar.Redis.Del(ctx, cookie).Result()

	duration := time.Since(start).Seconds()
	metrics.RedisQueryDuration.WithLabelValues("destroy_from_redis").Observe(duration)

	if err != nil {
		metrics.RedisQueryErrors.WithLabelValues("destroy_from_redis").Inc()
		errMsg := fmt.Errorf("redis: failed to destroy session. Error - %w", err)
		logger.Error().Err(errMsg).Msg("redis_destroy_error")
		return errMsg
	}

	logger.Info().Msg("redis: successfully destroy session")
	return nil
}
