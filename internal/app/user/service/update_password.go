package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) UpdatePassword(ctx context.Context, passwordData *models.PasswordData) *errVals.ServiceError {
	logger := log.Ctx(ctx)
	usr, err := u.userRepo.UserByID(ctx, passwordData.UserID)
	if err != nil {
		return errVals.ToServiceErrorFromRepo(err)
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(passwordData.OldPassword))
	if cryptErr != nil {
		logger.Err(cryptErr).Msg("BCrypt: password missmatched.")

		return errVals.NewServiceError(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidOldPassword)
	}

	err = u.userRepo.UpdatePassword(ctx, passwordData.UserID, passwordData.Password)

	if err != nil {
		return errVals.ToServiceErrorFromRepo(err)
	}

	return nil
}
