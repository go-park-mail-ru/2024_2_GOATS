package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	ck "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/user"
)

func (r *Repo) Session(ctx context.Context, cookie string) (*models.User, *errVals.ErrorObj, int) {
	cookieStore, err := ck.NewCookieStore(ctx)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	defer func() {
		if err := cookieStore.RedisDB.Close(); err != nil {
			log.Fatalf("Error closing output file %v", err)
		}
	}()

	userId, err := cookieStore.GetFromCookie(cookie)
	if err != nil || userId == "" {
		return nil, errVals.NewErrorObj(errVals.ErrUnauthorizedCode, errVals.CustomError{Err: err}), http.StatusUnauthorized
	}

	usr, err := user.FindById(ctx, userId, r.Database)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText), http.StatusNotFound
		}

		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return usr, nil, http.StatusOK
}
