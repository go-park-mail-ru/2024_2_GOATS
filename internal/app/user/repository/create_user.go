package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/user"
)

func (u *UserRepo) CreateUser(ctx context.Context, registerData *models.RegisterData) (*models.User, *errVals.ErrorObj, int) {
	hashPass, err := password.HashAndSalt(registerData.Password)
	if err != nil {
		return nil, &errVals.ErrorObj{
			Code: errVals.ErrServerCode,
			Error: errVals.CustomError{
				Err: fmt.Errorf("error hashing password: %w", err),
			},
		}, http.StatusInternalServerError
	}

	registerData.Password = hashPass

	usr, err := user.Create(ctx, *registerData, u.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrCreateUserCode, errVals.CustomError{Err: err}), http.StatusConflict
	}

	return usr, nil, http.StatusOK
}
