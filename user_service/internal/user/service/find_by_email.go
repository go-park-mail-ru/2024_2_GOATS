package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

// FindByEmail find user by email
func (u *UserService) FindByEmail(ctx context.Context, email string) (*dto.User, error) {
	logger := log.Ctx(ctx)
	repResp, err := u.userRepo.UserByEmail(ctx, email)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to get user by email")
		return nil, fmt.Errorf("userService - failed to get user by email: %w", err)
	}

	logger.Info().Msg("userService - successfully get user by email")
	return repResp, nil
}
