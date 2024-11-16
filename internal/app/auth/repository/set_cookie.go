package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func (r *AuthRepo) SetCookie(ctx context.Context, token *models.Token) (*models.CookieData, *errVals.RepoError) {
	logger := log.Ctx(ctx)
	cookieCfg := config.FromRedisContext(ctx).Cookie

	err := r.Redis.Set(ctx, token.TokenID, fmt.Sprint(token.UserID), cookieCfg.MaxAge)
	if err.Err() != nil {
		errMsg := fmt.Errorf("redis: cannot set cookie into redis - %w", err.Err())
		logger.Error().Err(errMsg).Msg("redis_set_error")

		return nil, errVals.NewRepoError(
			errVals.ErrRedisWriteCode,
			errVals.NewCustomError(errMsg.Error()),
		)
	}

	logger.Info().Msg("redis: successfully set cookie")

	return &models.CookieData{
		Name:  cookieCfg.Name,
		Token: token,
	}, nil
}
