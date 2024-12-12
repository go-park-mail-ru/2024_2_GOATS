package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

// CheckFavorite check user favorite
func (u *UserService) CheckFavorite(ctx context.Context, favData *dto.Favorite) (bool, error) {
	logger := log.Ctx(ctx)
	present, err := u.userRepo.CheckFavorite(ctx, converter.ConvertToRepoFavorite(favData))
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to checkFavorite")
		return false, fmt.Errorf("userService - failed to checkFavorite: %w", err)
	}

	logger.Info().Msg("userService - successfully checkFavorite")
	return present, nil
}
