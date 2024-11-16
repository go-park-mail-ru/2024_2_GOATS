package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/cookie"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, loginData *models.LoginData) (*models.AuthRespData, *errVals.ServiceError) {
	logger := log.Ctx(ctx)
	usr, err := s.userRepository.UserByEmail(ctx, loginData.Email)

	if err != nil {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(loginData.Password))
	if cryptErr != nil {
		logger.Err(cryptErr).Msg("BCrypt: password missmatched.")

		return nil, errVals.NewServiceError(errVals.ErrInvalidPasswordCode, errVals.NewCustomError(cryptErr.Error()))
	}

	token, tokenErr := cookie.GenerateToken(ctx, usr.ID)
	if tokenErr != nil {
		return nil, errVals.NewServiceError(errVals.ErrGenerateTokenCode, errVals.NewCustomError(tokenErr.Error()))
	}

	if loginData.Cookie != "" {
		err = s.authRepository.DestroySession(ctx, loginData.Cookie)
		if err != nil {
			return nil, errVals.ToServiceErrorFromRepo(err)
		}
	}

	ck, ckErr := s.authRepository.SetCookie(ctx, token)
	if ckErr != nil {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	return &models.AuthRespData{
		NewCookie: ck,
	}, nil
}
