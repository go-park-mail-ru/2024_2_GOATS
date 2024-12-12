package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

// GetSessionData get session data by cookie
func (as *AuthService) GetSessionData(ctx context.Context, cookie string) (uint64, error) {
	strUserID, err := as.authRepository.GetSessionData(ctx, cookie)
	if err != nil {
		return 0, fmt.Errorf("failed to getSessionData: %w", err)
	}

	usrID, err := strconv.ParseUint(strUserID, 10, 64)
	if err != nil {
		errMsg := fmt.Errorf("failed to getSessionData: failed to convert string into integer: %w", err)
		log.Ctx(ctx).Error().Err(errMsg).Msg("covertion_error")

		return 0, errMsg
	}

	return usrID, nil
}
