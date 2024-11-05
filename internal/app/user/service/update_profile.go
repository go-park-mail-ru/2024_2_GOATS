package service

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) UpdateProfile(ctx context.Context, usrData *models.User) (*models.UpdateUserRespData, *models.ErrorRespData) {
	if usrData.AvatarName != "" {
		avatarURL, err := u.userRepo.SaveUserAvatar(ctx, usrData)
		if err != nil {
			return nil, &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*err},
			}
		}
		usrData.AvatarURL = avatarURL
	}

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
