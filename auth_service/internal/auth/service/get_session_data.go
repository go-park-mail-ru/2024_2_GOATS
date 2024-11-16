package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errors"
	"github.com/rs/zerolog/log"
)

func (as *AuthService) GetSessionData(ctx context.Context, cookie string) (uint64, *errors.SrvErrorObj) {
	strUserID, err := as.authRepository.GetSessionData(ctx, cookie)
	if err != nil || strUserID == "" {
		return 0, nil
	}

	usrID, convErr := strconv.ParseUint(strUserID, 10, 64)
	if convErr != nil {
		errMsg := fmt.Errorf("session service: failed to convert string into integer: %w", convErr)
		log.Ctx(ctx).Error().Err(errMsg).Msg("covertion_error")
	}

	return usrID, nil
}
