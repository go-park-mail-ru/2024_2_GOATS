package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/subscriptiondb"
)

// CreateSubscription creates user subscription by calling db CreateSubscription
func (u *UserRepo) CreateSubscription(ctx context.Context, subData *dto.RepoCreateSubscriptionData) (uint64, error) {
	subID, err := subscriptiondb.CreateSubscription(ctx, subData, u.Database)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", "cannot_create_subscription", err)
	}

	return subID, nil
}
