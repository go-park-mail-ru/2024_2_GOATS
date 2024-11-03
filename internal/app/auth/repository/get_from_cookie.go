package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (r *Repo) GetFromCookie(ctx context.Context, cookie string) (string, *errVals.ErrorObj, int) {
	var userID string
	logger, requestId := config.FromBaseContext(ctx)

	err := r.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		errMsg := fmt.Errorf("redis: cannot get cookie from redis - %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return "", errVals.NewErrorObj(
			errVals.ErrCookieMissmatchCode,
			errVals.CustomError{Err: errMsg},
		), http.StatusForbidden
	}

	logger.Log("redis: successfully get info from cookie", requestId)
	return userID, nil, http.StatusOK
}
