package service

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/cookie"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *AuthService) Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *models.ErrorRespData) {
	usr, err, code := s.userRepository.CreateUser(ctx, registerData)
	if err != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorRespData{
			Errors:     errors,
			StatusCode: code,
		}
	}

	token, errVal := cookie.GenerateToken(ctx, usr.Id)
	if errVal != nil {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: errVal})},
			StatusCode: http.StatusInternalServerError,
		}
	}

	ck, errCk, code := s.authRepository.SetCookie(ctx, token)
	if errCk != nil {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*errCk},
			StatusCode: code,
		}
	}

	return &models.AuthRespData{
		NewCookie:  ck,
		StatusCode: code,
	}, nil
}
