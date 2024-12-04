package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

func (u *UserService) CreateSubscription(ctx context.Context, createData *dto.CreateSubscriptionData) (uint64, error) {
	logger := log.Ctx(ctx)
	repoData := converter.ConvertToRepoCreateSubData(createData)
	subID, err := u.userRepo.CreateSubscription(ctx, repoData)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to create subscription")
		return 0, fmt.Errorf("userService - failed to create subscription: %w", err)
	}

	return subID, nil
}
