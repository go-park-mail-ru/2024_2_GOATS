package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (s *AuthService) Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse) {
	userId, err, code := s.authRepository.GetFromCookie(ctx, cookie)
	if err != nil || userId == "" {
		return nil, &models.ErrorResponse{
			Success:    false,
			Errors:     []errVals.ErrorObj{*err},
			StatusCode: code,
		}
	}

	user, sesErr, code := s.authRepository.UserById(ctx, userId)
	if sesErr != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *sesErr

		return nil, &models.ErrorResponse{
			Success:    false,
			Errors:     errors,
			StatusCode: code,
		}
	}

	return &authModels.SessionResponse{
		Success:    true,
		StatusCode: code,
		UserData:   *user,
	}, nil
}
