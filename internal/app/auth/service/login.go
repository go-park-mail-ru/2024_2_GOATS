package service

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, loginData *models.LoginData) (*models.AuthRespData, *models.ErrorRespData) {
	usr, err, code := s.authRepository.UserByEmail(ctx, loginData)

	if err != nil {
		errors := make([]errors.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     errors,
		}
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(loginData.Password))
	if cryptErr != nil {
		return nil, &models.ErrorRespData{
			StatusCode: http.StatusConflict,
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText)},
		}
	}

	token, tokenErr := cookie.GenerateToken(ctx, usr.Id)
	if tokenErr != nil {
		return nil, &models.ErrorRespData{
			StatusCode: http.StatusInternalServerError,
			Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: tokenErr})},
		}
	}

	if loginData.Cookie != "" {
		err, code = s.authRepository.DestroySession(ctx, loginData.Cookie)
		if err != nil {
			return nil, &models.ErrorRespData{
				StatusCode: code,
				Errors:     []errVals.ErrorObj{*err},
			}
		}
	}

	ck, ckErr, code := s.authRepository.SetCookie(ctx, token)
	if ckErr != nil {
		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     []errVals.ErrorObj{*ckErr},
		}
	}

	return &models.AuthRespData{
		NewCookie:  ck,
		StatusCode: code,
	}, nil
}
