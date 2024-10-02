package service

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service/validation"
)

func (s *Service) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse) {
	success := true
	errors := make([]errVals.ErrorObj, 0)

	if err := validation.ValidatePassword(registerData.Password, registerData.PasswordConfirmation); err != nil {
		success = addError(errVals.ErrInvalidPasswordCode, *err, &errors)
	}

	if err := validation.ValidateEmail(registerData.Email); err != nil {
		success = addError(errVals.ErrInvalidEmailCode, *err, &errors)
	}

	if len(errors) > 0 {
		return nil, &models.ErrorResponse{
			Success:    success,
			Errors:     errors,
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	token, err, code := s.repository.Register(ctx, registerData)
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
		Token:   token,
		Success: true,
	}, nil
}

func addError(code string, err errVals.CustomError, errors *[]errVals.ErrorObj) bool {
	errStruct := errVals.ErrorObj{
		Code:  code,
		Error: err,
	}

	*errors = append(*errors, errStruct)

	return false
}
