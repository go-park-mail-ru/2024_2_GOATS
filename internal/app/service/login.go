package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (s *Service) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse) {
	cookies, err, code := s.repository.Login(ctx, loginData)

	if err != nil {
		errors := make([]errors.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: code,
			Errors:     errors,
		}
	}

	return &authModels.AuthResponse{
		ExpCookie:  cookies[0],
		NewCookie:  cookies[1],
		StatusCode: code,
		Success:    true,
	}, nil
}
