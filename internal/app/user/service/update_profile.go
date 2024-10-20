package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) UpdateProfile(ctx context.Context, profileData *models.User) (*models.UpdateUserRespData, *models.ErrorRespData) {
	err, status := u.userRepo.UpdateProfileData(ctx, profileData)
	if err != nil {
		return nil, &models.ErrorRespData{
			StatusCode: status,
			Errors:     []errVals.ErrorObj{*err},
		}
	}

	return &models.UpdateUserRespData{
		StatusCode: status,
	}, nil
}
