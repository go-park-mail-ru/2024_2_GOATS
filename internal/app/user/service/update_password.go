package service

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) UpdatePassword(ctx context.Context, passwordData *models.PasswordData) (*models.UpdateUserRespData, *models.ErrorRespData) {
	logger := log.Ctx(ctx)
	usr, err, status := u.userRepo.UserById(ctx, passwordData.UserId)
	if err != nil {
		return nil, &models.ErrorRespData{
			StatusCode: status,
			Errors:     []errVals.ErrorObj{*err},
		}
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(passwordData.OldPassword))
	if cryptErr != nil {
		logger.Err(cryptErr).Msg("BCrypt: password missmatch.")

		return nil, &models.ErrorRespData{
			StatusCode: http.StatusConflict,
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidOldPasswordText)},
		}
	}

	err, status = u.userRepo.UpdatePassword(ctx, passwordData.UserId, passwordData.Password)

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
