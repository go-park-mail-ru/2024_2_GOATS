package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

func (r *Repo) GetFromCookie(ctx context.Context, cookie string) (string, *errVals.ErrorObj, int) {
	var userID string
	logger := log.Ctx(ctx)

	err := r.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		errMsg := fmt.Errorf("redis: cannot get cookie from redis - %w", err)
		logger.Error().Msg(errMsg.Error())

		return "", errVals.NewErrorObj(
			errVals.ErrCreateUserCode,
			errVals.CustomError{Err: errMsg},
		), http.StatusForbidden
	}

	logger.Info().Msg(fmt.Sprintf("redis: successfully get info from cookie - %s", cookie))
	return userID, nil, http.StatusOK
}
