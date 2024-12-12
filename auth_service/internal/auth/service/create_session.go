package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/cookie"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
)

// CreateSession creates session by given params
func (as *AuthService) CreateSession(ctx context.Context, data *dto.SrvCreateCookie) (*dto.Cookie, error) {
	token, err := cookie.GenerateToken(ctx, data.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to createSession: %w", err)
	}

	ck, errCk := as.authRepository.SetCookie(ctx, converter.ToRepoTokenFromSrv(token))
	if errCk != nil {
		return nil, fmt.Errorf("failed to createSession: %w", errCk)
	}

	return ck, nil
}
