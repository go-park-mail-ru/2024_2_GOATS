package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (r *Repo) DestroySession(ctx context.Context, cookie string) (*errVals.ErrorObj, int) {
	_, err := r.Redis.Del(ctx, cookie).Result()

	if err != nil {
		return errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
