package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/user"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repo) Login(ctx context.Context, loginData *authModels.LoginData) ([]*authModels.CookieData, *errVals.ErrorObj, int) {
	usr, err := user.FindByEmail(ctx, loginData.Email, r.Database)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText), http.StatusNotFound
		}

		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(loginData.Password))
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText), http.StatusConflict
	}

	token, err := cookie.GenerateToken(ctx, usr.Id)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	var expCookie *authModels.CookieData

	cs := cookie.NewCookieStore(ctx, r.Redis)
	if loginData.Cookie != nil {
		expCookie, err = cs.DeleteCookie(loginData.Cookie.Value)
		if err != nil {
			return nil, errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
		}
	}

	ck, err := cs.SetCookie(token)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrRedisWriteCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	return []*authModels.CookieData{expCookie, ck}, nil, http.StatusOK
}
