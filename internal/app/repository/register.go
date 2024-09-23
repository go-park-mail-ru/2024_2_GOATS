package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/password"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (r *Repo) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.Token, *errVals.ErrorObj, int) {
	hashPass, err := password.HashAndSalt(registerData.Password)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	sqlStatement := `
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	usr := models.User{}
	err = r.Database.QueryRowContext(ctx, sqlStatement, registerData.Email, registerData.Username, hashPass).Scan(&usr.Id)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrCreateUserCode, errVals.CustomError{Err: err}), http.StatusConflict
	}

	token, err := cookie.GenerateToken(ctx, usr.Id)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	return token, nil, http.StatusOK
}
