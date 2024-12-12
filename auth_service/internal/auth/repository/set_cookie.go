package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/dto"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/metrics"
	"github.com/rs/zerolog/log"
)

// SetCookie sets cookie into Redis
func (ar *AuthRepository) SetCookie(ctx context.Context, token *dto.TokenData) (*srvDTO.Cookie, error) {
	start := time.Now()
	logger := log.Ctx(ctx)
	cookieCfg := config.FromRedisContext(ctx).Cookie

	err := ar.Redis.Set(ctx, token.TokenID, fmt.Sprint(token.UserID), cookieCfg.MaxAge)

	duration := time.Since(start).Seconds()
	metrics.RedisQueryDuration.WithLabelValues("insert_into_redis").Observe(duration)

	if err.Err() != nil {
		metrics.RedisQueryErrors.WithLabelValues("insert_into_redis").Inc()
		errMsg := fmt.Errorf("redis: cannot set cookie into redis - %w", err.Err())
		logger.Error().Err(errMsg).Msg("redis_set_error")

		return nil, errMsg
	}

	logger.Info().Msg("redis: successfully set cookie")

	return converter.ToCookieFromRepo(cookieCfg.Name, token), nil
}
