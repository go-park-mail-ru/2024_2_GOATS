package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/user"
)

func (u *UserRepo) UpdateProfileData(ctx context.Context, profileData *models.User) (*errVals.ErrorObj, int) {
	err := user.UpdateProfile(ctx, profileData, u.Database)
	if err != nil {
		return &errVals.ErrorObj{
			Code: "update_profile_error",
			Error: errVals.CustomError{
				Err: err,
			},
		}, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
