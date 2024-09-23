package repository

import (
	"context"
	"database/sql"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	ck "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/cookie"
)

func (r *Repo) Session(ctx context.Context, cookie string) (*models.User, *errVals.ErrorObj, int) {
	cookieStore, err := ck.NewCookieStore(ctx)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	defer cookieStore.RedisDB.Close()

	userId, err := cookieStore.GetFromCookie(ctx, cookie)
	if err != nil || userId == "" {
		return nil, errVals.NewErrorObj(errVals.ErrUnauthorizedCode, errVals.CustomError{Err: err}), http.StatusUnauthorized
	}

	var user models.User
	err = r.Database.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash FROM USERS WHERE id = $1", userId,
	).Scan(&user.Id, &user.Email, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText), http.StatusNotFound
		}
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return &user, nil, http.StatusOK
}
