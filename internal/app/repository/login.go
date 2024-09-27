package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/user"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repo) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.Token, *errVals.ErrorObj, int) {
	user, err := user.FindByEmail(ctx, loginData.Email, r.Database)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText), http.StatusNotFound
		}
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText), http.StatusConflict
	}

	token, err := cookie.GenerateToken(ctx, user.Id)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: err}), http.StatusInternalServerError
	}

	return token, nil, http.StatusOK
}
