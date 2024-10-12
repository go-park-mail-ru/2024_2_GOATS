package repository

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/repository/user"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (r *Repo) CreateUser(ctx context.Context, registerData *authModels.RegisterData) (*models.User, *errVals.ErrorObj, int) {
	usr, err := user.Create(ctx, *registerData, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrCreateUserCode, errVals.CustomError{Err: err}), http.StatusConflict
	}

	return usr, nil, http.StatusOK
}
