package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// GetFavorites gets user favorites
func (u *UserService) GetFavorites(ctx context.Context, usrID uint64) ([]uint64, error) {
	logger := log.Ctx(ctx)
	repResp, err := u.userRepo.GetFavorites(ctx, usrID)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to get user favorites")
		return nil, fmt.Errorf("userService - failed to get user favorites: %w", err)
	}

	logger.Info().Msg("userService - successfully get user favorites")
	return repResp, nil
}
