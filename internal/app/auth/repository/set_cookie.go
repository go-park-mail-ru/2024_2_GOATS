package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (r *Repo) SetCookie(ctx context.Context, token *models.Token) (*models.CookieData, *errVals.ErrorObj, int) {
	logger, requestId := config.FromBaseContext(ctx)
	cookieCfg := config.FromRedisContext(ctx).Cookie

	err := r.Redis.Set(ctx, token.TokenID, fmt.Sprint(token.UserID), cookieCfg.MaxAge)
	if err.Err() != nil {
		errMsg := fmt.Errorf("redis: cannot set cookie into redis - %w", err.Err())
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return nil, errVals.NewErrorObj(
			errVals.ErrCreateUserCode,
			errVals.CustomError{Err: errMsg},
		), http.StatusInternalServerError
	}

	logger.Log("redis: successfully set cookie", requestId)

	return &models.CookieData{
		Name:  cookieCfg.Name,
		Token: token,
	}, nil, http.StatusOK
}
