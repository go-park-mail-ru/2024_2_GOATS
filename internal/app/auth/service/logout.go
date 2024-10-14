package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *AuthService) Logout(ctx context.Context, cookie string) (*models.AuthRespData, *models.ErrorRespData) {
	err, code := s.authRepository.DestroySession(ctx, cookie)

	if err != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorRespData{
			Errors:     errors,
			StatusCode: code,
		}
	}

	return &models.AuthRespData{
		StatusCode: code,
	}, nil
}
