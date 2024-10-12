package service

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service/validation"
)

func (s *Service) Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse) {
	validationErr := validation.ValidateCookie(cookie)
	if validationErr != nil {
		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: http.StatusForbidden,
			Errors:     []errVals.ErrorObj{{Code: errVals.ErrBrokenCookieCode, Error: *validationErr}},
		}
	}

	user, err, code := s.repository.Session(ctx, cookie)
	if err != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *err

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
