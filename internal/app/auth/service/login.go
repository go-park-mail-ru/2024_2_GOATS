package service

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse) {
	usr, err, code := s.authRepository.UserByEmail(ctx, loginData)

	if err != nil {
		errors := make([]errors.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: code,
			Errors:     errors,
		}
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(loginData.Password))
	if cryptErr != nil {
		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: http.StatusConflict,
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText)},
		}
	}

	token, tokenErr := cookie.GenerateToken(ctx, usr.Id)
	if tokenErr != nil {
		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: tokenErr})},
		}
	}

	if loginData.Cookie != "" {
		err, code = s.authRepository.DestroySession(ctx, loginData.Cookie)
		if err != nil {
			return nil, &models.ErrorResponse{
				Success:    false,
				StatusCode: code,
				Errors:     []errVals.ErrorObj{*err},
			}
		}
	}

	ck, ckErr, code := s.authRepository.SetCookie(ctx, token)
	if ckErr != nil {
		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: code,
			Errors:     []errVals.ErrorObj{*ckErr},
		}
	}

	return &authModels.AuthResponse{
		NewCookie:  ck,
		StatusCode: code,
		Success:    true,
	}, nil
}
