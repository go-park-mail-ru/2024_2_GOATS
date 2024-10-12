package service

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/validation"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (s *AuthService) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse) {
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

	usr, err, code := s.authRepository.CreateUser(ctx, registerData)
	if err != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorResponse{
			Success:    false,
			Errors:     errors,
			StatusCode: code,
		}
	}

	token, errVal := cookie.GenerateToken(ctx, usr.Id)
	if errVal != nil {
		return nil, &models.ErrorResponse{
			Success:    false,
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: errVal})},
			StatusCode: http.StatusInternalServerError,
		}
	}

	ck, errCk, code := s.authRepository.SetCookie(ctx, token)
	if errCk != nil {
		return nil, &models.ErrorResponse{
			Success:    false,
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisWriteCode, errVals.CustomError{Err: errCk})},
			StatusCode: code,
		}
	}

	return &authModels.AuthResponse{
		NewCookie:  ck,
		StatusCode: code,
		Success:    true,
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
