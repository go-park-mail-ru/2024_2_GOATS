package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (r *Repo) GetFromCookie(ctx context.Context, cookie string) (string, *errVals.ErrorObj, int) {
	var userID string
	err := r.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		return "", errVals.NewErrorObj(
			errVals.ErrCreateUserCode,
			errVals.CustomError{Err: fmt.Errorf("cannot get cookie from redis: %w", err)},
		), http.StatusForbidden
	}

	return userID, nil, http.StatusOK
}
