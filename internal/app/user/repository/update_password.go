package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) UpdatePassword(ctx context.Context, usrID int, pass string) (*errVals.ErrorObj, int) {
	hashedPasswd, err := password.HashAndSalt(ctx, pass)
	if err != nil {
		return errVals.NewErrorObj(
			errVals.ErrServerCode,
			errVals.CustomError{Err: fmt.Errorf("error hashing password: %w", err)},
		), http.StatusInternalServerError
	}

	err = userdb.UpdatePassword(ctx, usrID, hashedPasswd, u.Database)
	if err != nil {
		return errVals.NewErrorObj(
			errVals.ErrUpdatePasswordCode,
			errVals.CustomError{Err: fmt.Errorf("error updating password: %w", err)},
		), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
