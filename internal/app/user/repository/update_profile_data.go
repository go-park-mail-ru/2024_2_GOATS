package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) UpdateProfileData(ctx context.Context, profileData *models.User) (*errVals.ErrorObj, int) {
	err := userdb.UpdateProfile(ctx, profileData, u.Database)
	if err != nil {
		return errVals.NewErrorObj(
			errVals.ErrUpdateProfileCode,
			errVals.CustomError{Err: fmt.Errorf("error updating profile: %w", err)},
		), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
