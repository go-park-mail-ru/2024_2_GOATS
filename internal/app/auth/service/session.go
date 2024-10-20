package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *AuthService) Session(ctx context.Context, cookie string) (*models.SessionRespData, *models.ErrorRespData) {
	userId, err, code := s.authRepository.GetFromCookie(ctx, cookie)
	if err != nil || userId == "" {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*err},
			StatusCode: code,
		}
	}

	user, sesErr, code := s.authRepository.UserById(ctx, userId)
	if sesErr != nil {
		errs := make([]errVals.ErrorObj, 1)
		errs[0] = *sesErr

		return nil, &models.ErrorRespData{
			Errors:     errs,
			StatusCode: code,
		}
	}

	return &models.SessionRespData{
		StatusCode: code,
		UserData:   *user,
	}, nil
}
