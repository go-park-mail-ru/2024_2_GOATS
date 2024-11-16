package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errors"
	"github.com/rs/zerolog/log"
)

func (ar *AuthRepository) DestroySession(ctx context.Context, cookie string) *errors.ErrorObj {
	logger := log.Ctx(ctx)
	_, err := ar.Redis.Del(ctx, cookie).Result()

	if err != nil {
		errMsg := fmt.Errorf("redis: failed to destroy session. Error - %w", err)
		logger.Error().Err(errMsg).Msg("redis_destroy_error")
		return nil
	}

	logger.Info().Msg("redis: successfully destroy session")
	return nil
}
