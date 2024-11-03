package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (r *Repo) DestroySession(ctx context.Context, cookie string) (*errVals.ErrorObj, int) {
	logger, requestId := config.FromBaseContext(ctx)
	_, err := r.Redis.Del(ctx, cookie).Result()

	if err != nil {
		errMsg := fmt.Errorf("redis: failed to destroy session. Error - %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)
		return errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errMsg}), http.StatusInternalServerError
	}

	logger.Log("redis: successfully destroy session", requestId)
	return nil, http.StatusOK
}
