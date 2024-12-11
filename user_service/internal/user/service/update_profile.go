package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

// UpdateProfile saves avatar and updates user profile by calling userRepo UpdateProfile
func (u *UserService) UpdateProfile(ctx context.Context, usrData *dto.User) error {
	logger := log.Ctx(ctx)
	if usrData.AvatarName != "" {
		avatarURL, err := u.userRepo.SaveUserAvatar(ctx, usrData.AvatarName, usrData.AvatarFile)

		if err != nil {
			logger.Error().Err(err).Msg("userService - failed to updateProfile")
			return fmt.Errorf("userService - failed to updateProfile: %w", err)
		}

		usrData.AvatarURL = avatarURL
	}

	err := u.userRepo.UpdateProfileData(ctx, converter.ConvertToRepoUser(usrData))
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to updateProfile")
		return fmt.Errorf("userService failed to updateProfile: %w", err)
	}

	return nil
}
