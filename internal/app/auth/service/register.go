package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/cookie"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *AuthService) Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *errVals.ServiceError) {
	usr, err := s.userRepository.CreateUser(ctx, converter.ToDBRegisterFromRegister(registerData))
	if err != nil {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	token, errVal := cookie.GenerateToken(ctx, usr.ID)
	if errVal != nil {
		return nil, errVals.NewServiceError(errVals.ErrGenerateTokenCode, errVals.NewCustomError(errVal.Error()))
	}

	ck, errCk := s.authRepository.SetCookie(ctx, token)
	if errCk != nil {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	return &models.AuthRespData{
		NewCookie: ck,
	}, nil
}
