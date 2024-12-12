package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

// SetFavorite sets user favorite by calling userRepo SetFavorite
func (u *UserService) SetFavorite(ctx context.Context, favData *dto.Favorite) error {
	logger := log.Ctx(ctx)
	err := u.userRepo.SetFavorite(ctx, converter.ConvertToRepoFavorite(favData))
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to setFavorite")
		return fmt.Errorf("userService - failed to setFavorite: %w", err)
	}

	logger.Info().Msg("userService - successfully setFavorite")
	return nil
}
