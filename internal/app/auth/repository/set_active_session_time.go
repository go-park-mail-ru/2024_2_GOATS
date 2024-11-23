package repository

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (r *AuthRepo) SetActiveSessionTime(ctx context.Context, usrID string, seconds int) error {
	logger := log.Ctx(ctx)

	err := r.Redis.Set(ctx, usrID, seconds, 0)
	if err.Err() != nil {
		errMsg := fmt.Errorf("redis: cannot set active time into redis - %w", err.Err())
		logger.Error().Err(errMsg).Msg("redis_set_error")

		return errMsg
	}

	fmt.Println(usrID)
	fmt.Println(seconds)
	logger.Info().Msg("redis: successfully set active time")

	return nil
}
