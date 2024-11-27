package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) UpdatePassword(ctx context.Context, passwordData *dto.PasswordData) error {
	logger := log.Ctx(ctx)
	usr, err := u.userRepo.UserByID(ctx, passwordData.UserID)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to update password")
		return fmt.Errorf("userService - failed to update password: %w", err)
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(passwordData.OldPassword))
	if cryptErr != nil {
		logger.Error().Err(cryptErr).Msg("BCrypt: password missmatched.")
		return fmt.Errorf("userService failed to update password: %s: %w", errVals.ErrInvalidPasswordCode, errVals.ErrInvalidOldPassword)
	}

	err = u.userRepo.UpdatePassword(ctx, passwordData.UserID, passwordData.Password)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to update password")
		return fmt.Errorf("userService failed to update password: %w", err)
	}

	return nil
}
