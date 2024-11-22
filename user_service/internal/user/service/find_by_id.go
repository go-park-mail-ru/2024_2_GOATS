package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

func (u *UserService) FindByID(ctx context.Context, usrID uint64) (*dto.User, error) {
	logger := log.Ctx(ctx)
	repResp, err := u.userRepo.UserByID(ctx, usrID)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to get user by id")
		return nil, fmt.Errorf("userService - failed to get user by id: %w", err)
	}

	logger.Info().Msg("userService - successfully get user by id")
	return repResp, nil
}
