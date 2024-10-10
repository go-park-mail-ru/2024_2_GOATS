package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/user"
)

func (r *Repo) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.CookieData, *errVals.ErrorObj, int) {
	usr, err := user.Create(ctx, *registerData, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrCreateUserCode, errVals.CustomError{Err: err}), http.StatusConflict
	}

	token, err := cookie.GenerateToken(ctx, usr.Id)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	cs := cookie.NewCookieStore(ctx, r.Redis)
	ck, err := cs.SetCookie(token)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrRedisWriteCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	return ck, nil, http.StatusOK
}
