package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errors"
	"github.com/rs/zerolog/log"
)

func (ar *AuthRepository) GetSessionData(ctx context.Context, cookie string) (string, *errors.ErrorObj) {
	var userID string
	logger := log.Ctx(ctx)

	err := ar.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		errMsg := fmt.Errorf("redis: cannot get cookie from redis - %w", err)
		logger.Error().Err(errMsg).Msg("redis_get_error")

		return "", nil
	}

	logger.Info().Msg("redis: successfully get info from cookie")
	return userID, nil
}
