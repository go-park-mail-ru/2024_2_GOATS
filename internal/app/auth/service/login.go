package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, loginData *models.LoginData) (*models.AuthRespData, *errVals.ServiceError) {
	logger := log.Ctx(ctx)
	usr, err := s.userClient.FindByEmail(ctx, loginData.Email)

	if err != nil {
		return nil, errVals.NewServiceError(errVals.ErrGetUserCode, fmt.Errorf("failed to login: %w", err))
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(loginData.Password))
	if cryptErr != nil {
		logger.Err(cryptErr).Msg("BCrypt: password missmatched.")

		return nil, errVals.NewServiceError(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatch.Err)
	}

	if loginData.Cookie != "" {
		err = s.authClient.DestroySession(ctx, loginData.Cookie)
		if err != nil {
			return nil, errVals.NewServiceError(errVals.ErrDestroySessionCode, fmt.Errorf("failed to login: %w", err))
		}
	}

	ckResp, err := s.authClient.CreateSession(ctx, usr.ID)
	if err != nil {
		return nil, errVals.NewServiceError(errVals.ErrCreateSessionCode, fmt.Errorf("failed to login: %w", err))
	}

	return &models.AuthRespData{
		NewCookie: &models.CookieData{
			Name:  ckResp.Name,
			Token: ckResp.Token,
		},
	}, nil
}
