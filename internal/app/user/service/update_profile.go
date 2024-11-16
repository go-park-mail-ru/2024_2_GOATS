package service

import (
	"context"
	"io"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/converter"
)

func (u *UserService) UpdateProfile(ctx context.Context, usrData *models.User) *errVals.ServiceError {
	if usrData.AvatarName != "" {
		avatarURL, avatarFile, err := u.userRepo.SaveUserAvatar(ctx, usrData.AvatarName)
		if err != nil {
			return errVals.ToServiceErrorFromRepo(err)
		}
		_, fileErr := io.Copy(avatarFile, usrData.AvatarFile)
		if fileErr != nil {
			return errVals.NewServiceError(errVals.ErrSaveFileCode, errVals.ErrSaveFile)
		}

		usrData.AvatarURL = avatarURL
	}

	err := u.userRepo.UpdateProfileData(ctx, converter.ToDBUserFromUser(usrData))
	if err != nil {
		return errVals.ToServiceErrorFromRepo(err)
	}

	return nil
}
