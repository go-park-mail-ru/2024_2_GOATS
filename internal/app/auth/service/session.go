package service

import (
	"context"
	"fmt"
	"strconv"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func (s *AuthService) Session(ctx context.Context, cookie string) (*models.SessionRespData, *models.ErrorRespData) {
	strUserID, err, code := s.authRepository.GetFromCookie(ctx, cookie)
	if err != nil || strUserID == "" {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*err},
			StatusCode: code,
		}
	}

	usrID, convErr := strconv.Atoi(strUserID)
	if convErr != nil {
		errMsg := fmt.Errorf("session service: failed to convert string into integer: %w", convErr)
		log.Ctx(ctx).Error().Err(errMsg).Msg("covertion_error")

		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrConvertionCode, errVals.CustomError{Err: errMsg})},
			StatusCode: code,
		}
	}

	user, sesErr, code := s.userRepository.UserByID(ctx, usrID)
	if sesErr != nil {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*sesErr},
			StatusCode: code,
		}
	}

	return &models.SessionRespData{
		StatusCode: code,
		UserData:   *user,
	}, nil
}
