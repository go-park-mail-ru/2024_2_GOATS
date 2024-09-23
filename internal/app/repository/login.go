package repository

import (
	"context"
	"database/sql"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/cookie"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repo) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.Token, *errVals.ErrorObj, int) {
	var user models.User
	err := r.Database.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash FROM USERS WHERE email = $1", loginData.Email,
	).Scan(&user.Id, &user.Email, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
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
