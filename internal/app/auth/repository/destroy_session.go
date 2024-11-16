package repository

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

func (r *AuthRepo) DestroySession(ctx context.Context, cookie string) *errVals.RepoError {
	logger := log.Ctx(ctx)
	_, err := r.Redis.Del(ctx, cookie).Result()

	if err != nil {
		errMsg := fmt.Errorf("redis: failed to destroy session. Error - %w", err)
		logger.Error().Err(errMsg).Msg("redis_destroy_error")
		return errVals.NewRepoError(errVals.ErrRedisClearCode, errVals.NewCustomError(errMsg.Error()))
	}

	logger.Info().Msg("redis: successfully destroy session")
	return nil
}
