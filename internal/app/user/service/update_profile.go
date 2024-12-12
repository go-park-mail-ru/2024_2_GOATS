package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// UpdateProfile updates user profile by calling userClient UpdateProfile
func (u *UserService) UpdateProfile(ctx context.Context, usrData *models.User) *errVals.ServiceError {
	err := u.userClient.UpdateProfile(ctx, usrData)
	if err != nil {
		return errVals.NewServiceError("update_profile_error", err)
	}

	return nil
}
