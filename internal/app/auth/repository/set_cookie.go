package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) SetCookie(ctx context.Context, token *models.Token) (*models.CookieData, *errVals.ErrorObj, int) {
	logger := log.Ctx(ctx)
	cookieCfg := config.FromRedisContext(ctx).Cookie

	err := r.Redis.Set(ctx, token.TokenID, fmt.Sprint(token.UserID), cookieCfg.MaxAge)
	if err.Err() != nil {
		errMsg := fmt.Errorf("redis: cannot set cookie into redis - %w", err.Err())
		logger.Error().Msg(errMsg.Error())

		return nil, errVals.NewErrorObj(
			errVals.ErrCreateUserCode,
			errVals.CustomError{Err: errMsg},
		), http.StatusInternalServerError
	}

	logger.Info().Msg(fmt.Sprintf("redis: successfully set cookie - %s", token.TokenID))

	return &models.CookieData{
		Name:  cookieCfg.Name,
		Token: token,
	}, nil, http.StatusOK
}
