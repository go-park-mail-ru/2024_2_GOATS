package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) UserByID(ctx context.Context, userID int) (*models.User, *errVals.ErrorObj, int) {
	usr, err := userdb.FindByID(ctx, userID, u.Database)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText), http.StatusNotFound
		}

		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return usr, nil, http.StatusOK
}
