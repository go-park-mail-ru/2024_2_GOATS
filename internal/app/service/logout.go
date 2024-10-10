package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (s *Service) Logout(ctx context.Context, cookie string) (*auth.AuthResponse, *models.ErrorResponse) {
	expCookie, err, code := s.repository.Logout(ctx, cookie)

	if err != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorResponse{
			Success:    false,
			Errors:     errors,
			StatusCode: code,
		}
	}

	return &authModels.AuthResponse{
		Success:    true,
		StatusCode: code,
		ExpCookie:  expCookie,
	}, nil
}
