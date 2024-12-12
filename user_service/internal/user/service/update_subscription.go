package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// UpdateSubscribtionStatus updates user's subscriprion status
func (u *UserService) UpdateSubscribtionStatus(ctx context.Context, subID uint64) error {
	logger := log.Ctx(ctx)
	err := u.userRepo.UpdateSubscribtionStatus(ctx, subID)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to update subscription status")
		return fmt.Errorf("userService - failed to update subscription status: %w", err)
	}

	return nil
}
