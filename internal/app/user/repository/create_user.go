package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) CreateUser(ctx context.Context, registerData *models.RegisterData) (*models.User, *errVals.ErrorObj, int) {
	hashedPasswd, err := password.HashAndSalt(ctx, registerData.Password)
	if err != nil {
		return nil, errVals.NewErrorObj(
			errVals.ErrServerCode,
			errVals.CustomError{
				Err: fmt.Errorf("error hashing password: %w", err),
			}), http.StatusInternalServerError
	}

	registerData.Password = hashedPasswd

	usr, err := userdb.Create(ctx, *registerData, u.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrCreateUserCode, errVals.CustomError{Err: err}), http.StatusConflict
	}

	return usr, nil, http.StatusOK
}
