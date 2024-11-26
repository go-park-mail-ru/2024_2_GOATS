package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/converter"
	"github.com/rs/zerolog/log"
)

func (u *UserService) UpdateProfile(ctx context.Context, usrData *models.User) *errVals.ServiceError {
	err := u.userClient.UpdateProfile(ctx, usrData)
	if err != nil {
		return errVals.NewServiceError("update_profile_error", err)
	}

	return nil
}
