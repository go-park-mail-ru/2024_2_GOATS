package service

import (
	"context"
	"fmt"
	"io"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/converter"
	"github.com/rs/zerolog/log"
)

func (u *UserService) UpdateProfile(ctx context.Context, usrData *models.User) *errVals.ServiceError {
	if usrData.AvatarName != "" {
		avatarURL, avatarFile, err := u.userRepo.SaveUserAvatar(ctx, usrData.AvatarName)
		defer func() {
			if err := avatarFile.Close(); err != nil {
				log.Ctx(ctx).Err(fmt.Errorf("failed to close outFile: %w", err))
			}
		}()

		if err != nil {
			return errVals.ToServiceErrorFromRepo(err)
		}
		_, fileErr := io.Copy(avatarFile, usrData.AvatarFile)
		if fileErr != nil {
			return errVals.NewServiceError(errVals.ErrSaveFileCode, errVals.ErrSaveFile)
		}

		usrData.AvatarURL = avatarURL
	}

	err := u.userRepo.UpdateProfileData(ctx, converter.ToRepoUserFromUser(usrData))
	if err != nil {
		return errVals.ToServiceErrorFromRepo(err)
	}

	return nil
}
