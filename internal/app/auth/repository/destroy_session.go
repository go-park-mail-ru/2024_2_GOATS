package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

func (r *Repo) DestroySession(ctx context.Context, cookie string) (*errVals.ErrorObj, int) {
	logger := log.Ctx(ctx)
	_, err := r.Redis.Del(ctx, cookie).Result()

	if err != nil {
		errMsg := fmt.Errorf("redis: failed to destroy session. Error - %w", err)
		logger.Error().Msg(errMsg.Error())
		return errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errMsg}), http.StatusInternalServerError
	}

	logger.Info().Msg("redis: successfully destroy session")
	return nil, http.StatusOK
}
