package repository

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

func (r *AuthRepo) GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError) {
	var userID string
	logger := log.Ctx(ctx)

	err := r.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		errMsg := fmt.Errorf("redis: cannot get cookie from redis - %w", err)
		logger.Error().Err(errMsg).Msg("redis_get_error")

		return "", errVals.NewRepoError(
			errVals.ErrRedisGetCode,
			errVals.NewCustomError(errMsg.Error()),
		)
	}

	logger.Info().Msg("redis: successfully get info from cookie")
	return userID, nil
}
