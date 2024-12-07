package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/subscriptiondb"
)

func (u *UserRepo) UpdateSubscribtionStatus(ctx context.Context, subID uint64) error {
	err := subscriptiondb.UpdateSubscription(ctx, subID, u.Database)
	if err != nil {
		return fmt.Errorf("%s: %w", "cannot_update_subscription", err)
	}

	return nil
}
