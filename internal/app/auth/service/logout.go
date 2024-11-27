package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (s *AuthService) Logout(ctx context.Context, cookie string) *errVals.ServiceError {
	err := s.authClient.DestroySession(ctx, cookie)

	if err != nil {
		return errVals.NewServiceError(errVals.ErrDestroySessionCode, fmt.Errorf("failed to logout: %w", err))
	}

	return nil
}
