package service

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) UpdateProfile(ctx context.Context, usrData *models.User) (*models.UpdateUserRespData, *models.ErrorRespData) {
	avatarUrl, err := u.userRepo.SaveAvatar(ctx, usrData)
	if err != nil {
		return nil, &models.ErrorRespData{
			StatusCode: http.StatusInternalServerError,
			Errors:     []errVals.ErrorObj{*err},
		}
	}

	usrData.AvatarUrl = avatarUrl
	err, status := u.userRepo.UpdateProfileData(ctx, usrData)
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
