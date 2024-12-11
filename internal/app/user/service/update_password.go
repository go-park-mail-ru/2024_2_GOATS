package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// UpdatePassword updates user password by calling userClient methods
func (u *UserService) UpdatePassword(ctx context.Context, passwordData *models.PasswordData) *errVals.ServiceError {
	logger := log.Ctx(ctx)
	usr, err := u.userClient.FindByID(ctx, uint64(passwordData.UserID))
	if err != nil {
		return errVals.NewServiceError("failed_to_update_password", err)
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(passwordData.OldPassword))
	if cryptErr != nil {
		logger.Err(cryptErr).Msg("BCrypt: password missmatched.")
		return errVals.NewServiceError("failed_to_update_password", errVals.ErrInvalidOldPassword.Err)
	}

	err = u.userClient.UpdatePassword(ctx, passwordData)
	if err != nil {
		return errVals.NewServiceError("failed_to_update_password", err)
	}

	return nil
}
