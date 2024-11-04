package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *AuthService) Logout(ctx context.Context, cookie string) (*models.AuthRespData, *models.ErrorRespData) {
	err, code := s.authRepository.DestroySession(ctx, cookie)

	if err != nil {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*err},
			StatusCode: code,
		}
	}

	return &models.AuthRespData{
		StatusCode: code,
	}, nil
}
