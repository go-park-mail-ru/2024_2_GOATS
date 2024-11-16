package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errors"
)

func (as *AuthService) CreateSession(ctx context.Context, data *dto.SrvCreateCookie) (*dto.Cookie, *errors.SrvErrorObj) {
	token, errVal := cookie.GenerateToken(ctx, data.UserID)
	if errVal != nil {
		return nil, nil
	}

	ck, errCk := as.authRepository.SetCookie(ctx, converter.ToRepoTokenFromSrv(token))
	if errCk != nil {
		return nil, nil
	}

	return ck, nil
}
