package subscriptiondb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	metricsutils "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/metrics_utils"
	"github.com/rs/zerolog/log"
)

const (
	subCreateSQL = `
		INSERT INTO subscriptions (user_id, price, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	markPaidSQL = "UPDATE subscriptions SET status = $1, expiration_date = $2 WHERE id = $3"

	pendingStatus = "pending"
	activeStatus  = "active"
)

func CreateSubscription(ctx context.Context, subData *dto.RepoCreateSubscriptionData, db *sql.DB) (uint64, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	var subID uint64
	err := db.QueryRowContext(
		ctx,
		subCreateSQL,
		subData.UserID, subData.Amount, pendingStatus,
	).Scan(&subID)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "create_subscription", "subscriptions")
		errMsg := fmt.Errorf("postgres: error while creating subscription - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return 0, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "create_subscription", "subscriptions")
	logger.Info().Msg("postgres: subscription created successfully")

	return subID, nil
}

func UpdateSubscription(ctx context.Context, subID uint64, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	_, err := db.ExecContext(ctx, markPaidSQL, activeStatus, time.Now().AddDate(0, 1, 0), subID)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "update_subscription_status", "subscriptions")
		errMsg := fmt.Errorf("postgres: error while updating subscription status - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "update_subscription_status", "subscriptions")
	logger.Info().Msg("postgres: subscription status updated successfully")

	return nil
}
