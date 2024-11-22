package repository

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (ar *AuthRepository) DestroySession(ctx context.Context, cookie string) error {
	logger := log.Ctx(ctx)
	_, err := ar.Redis.Del(ctx, cookie).Result()

	if err != nil {
		errMsg := fmt.Errorf("redis: failed to destroy session. Error - %w", err)
		logger.Error().Err(errMsg).Msg("redis_destroy_error")
		return errMsg
	}

	logger.Info().Msg("redis: successfully destroy session")
	return nil
}
