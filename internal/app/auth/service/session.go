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
	strUserId, err, code := s.authRepository.GetFromCookie(ctx, cookie)
	if err != nil || strUserId == "" {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*err},
			StatusCode: code,
		}
	}

	usrId, convErr := strconv.Atoi(strUserId)
	if convErr != nil {
		errMsg := fmt.Errorf("Session service: failed to convert string into integer: %w", convErr)
		log.Ctx(ctx).Error().Msg(errMsg.Error())

		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj("convertion_error", errVals.CustomError{Err: errMsg})},
			StatusCode: code,
		}
	}

	user, sesErr, code := s.userRepository.UserById(ctx, usrId)
	if sesErr != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *sesErr

		return nil, &models.ErrorRespData{
			Errors:     errors,
			StatusCode: code,
		}
	}

	return &models.SessionRespData{
		StatusCode: code,
		UserData:   *user,
	}, nil
}
