package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/user"
)

func (u *UserRepo) UpdatePassword(ctx context.Context, usrId int, pass string) (*errVals.ErrorObj, int) {
	hashPass, err := password.HashAndSalt(pass)
	if err != nil {
		return &errVals.ErrorObj{
			Code: errVals.ErrServerCode,
			Error: errVals.CustomError{
				Err: fmt.Errorf("error hashing password: %w", err),
			},
		}, http.StatusInternalServerError
	}

	err = user.UpdatePassword(ctx, usrId, hashPass, u.Database)
	if err != nil {
		return &errVals.ErrorObj{
			Code: errVals.ErrUpdatePasswordCode,
			Error: errVals.CustomError{
				Err: err,
			},
		}, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
