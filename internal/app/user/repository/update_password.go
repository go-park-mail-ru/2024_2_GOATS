package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/user"
)

func (u *UserRepo) UpdatePassword(ctx context.Context, usrId int, password string) (*errVals.ErrorObj, int) {
	err := user.UpdatePassword(ctx, usrId, password, u.Database)
	if err != nil {
		return &errVals.ErrorObj{
			Code: "update_password_error",
			Error: errVals.CustomError{
				Err: err,
			},
		}, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
