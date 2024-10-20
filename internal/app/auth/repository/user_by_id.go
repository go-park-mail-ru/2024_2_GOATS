package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/repository/user"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (r *Repo) UserById(ctx context.Context, userId string) (*models.User, *errVals.ErrorObj, int) {
	usr, err := user.FindById(ctx, userId, r.Database)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText), http.StatusNotFound
		}

		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return usr, nil, http.StatusOK
}
