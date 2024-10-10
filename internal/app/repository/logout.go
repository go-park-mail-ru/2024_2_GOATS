package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	ck "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/cookie"
	"github.com/labstack/gommon/log"
)

func (r *Repo) Logout(ctx context.Context, cookie string) (*authModels.CookieData, *errVals.ErrorObj, int) {
	expCookie, err := ck.NewCookieStore(ctx, r.Redis).DeleteCookie(cookie)
	if err != nil {
		log.Errorf("cookie error: %v", err)
		return nil, errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	return expCookie, nil, http.StatusOK
}
