package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errors"
)

func (as *AuthService) DestroySession(ctx context.Context, cookie string) (bool, *errors.SrvErrorObj) {
	err := as.authRepository.DestroySession(ctx, cookie)

	if err != nil {
		return false, nil
	}

	return true, nil
}
