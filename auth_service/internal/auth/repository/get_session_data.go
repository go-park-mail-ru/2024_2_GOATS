package repository

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (ar *AuthRepository) GetSessionData(ctx context.Context, cookie string) (string, error) {
	var userID string
	logger := log.Ctx(ctx)

	err := ar.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		errMsg := fmt.Errorf("redis: cannot get cookie from redis - %w", err)
		logger.Error().Err(errMsg).Msg("redis_get_error")

		return "", errMsg
	}

	logger.Info().Msg("redis: successfully get info from cookie")
	return userID, nil
}
