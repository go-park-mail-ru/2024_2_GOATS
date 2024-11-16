package service

import (
	"context"
	"fmt"
	"strconv"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func (s *AuthService) Session(ctx context.Context, cookie string) (*models.SessionRespData, *errVals.ServiceError) {
	strUserID, err := s.authRepository.GetFromCookie(ctx, cookie)
	if err != nil || strUserID == "" {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	usrID, convErr := strconv.Atoi(strUserID)
	if convErr != nil {
		errMsg := fmt.Errorf("session service: failed to convert string into integer: %w", convErr)
		log.Ctx(ctx).Error().Err(errMsg).Msg("covertion_error")

		return nil, errVals.NewServiceError(errVals.ErrConvertionCode, errVals.NewCustomError(errMsg.Error()))
	}

	user, sesErr := s.userRepository.UserByID(ctx, usrID)
	if sesErr != nil {
		return nil, errVals.ToServiceErrorFromRepo(sesErr)
	}

	return &models.SessionRespData{
		UserData: *user,
	}, nil
}
