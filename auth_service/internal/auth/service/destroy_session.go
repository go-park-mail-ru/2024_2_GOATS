package service

import (
	"context"
	"fmt"
)

func (as *AuthService) DestroySession(ctx context.Context, cookie string) (bool, error) {
	err := as.authRepository.DestroySession(ctx, cookie)

	if err != nil {
		return false, fmt.Errorf("failed to destroySession: %w", err)
	}

	return true, nil
}
