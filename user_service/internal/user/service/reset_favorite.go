package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

func (u *UserService) ResetFavorite(ctx context.Context, favData *dto.Favorite) error {
	logger := log.Ctx(ctx)
	err := u.userRepo.ResetFavorite(ctx, converter.ConvertToRepoFavorite(favData))
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to reset favorite")
		return fmt.Errorf("userService failed to reset favorite: %w", err)
	}

	logger.Info().Msg("userService - successfully reset favorite")
	return nil
}
