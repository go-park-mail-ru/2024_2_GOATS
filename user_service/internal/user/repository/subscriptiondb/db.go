package subscriptiondb

import (
	"context"
	"database/sql"
	"errors"
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

	markPaidSQL     = "UPDATE subscriptions SET status = $1, expiration_date = $2 WHERE id = $3"
	findByUserIDSQL = "SELECT status, expiration_date FROM subscriptions WHERE user_id = $1"

	PendingStatus = "pending"
	ActiveStatus  = "active"
)

func CreateSubscription(ctx context.Context, subData *dto.RepoCreateSubscriptionData, db *sql.DB) (uint64, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	var subID uint64
	err := db.QueryRowContext(
		ctx,
		subCreateSQL,
		subData.UserID, subData.Amount, PendingStatus,
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

	_, err := db.ExecContext(ctx, markPaidSQL, ActiveStatus, time.Now().AddDate(0, 1, 0), subID)

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

func FindByUserID(ctx context.Context, usrID uint64, db *sql.DB) (*dto.RepoSubscription, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	var sub = &dto.RepoSubscription{}
	row := db.QueryRowContext(ctx, findByUserIDSQL, usrID)

	err := row.Scan(
		&sub.Status,
		&sub.ExpirationDate,
	)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			metricsutils.SaveErrorMetric(start, "get_actor_by_id", "actors")
			errMsg := fmt.Errorf("postgres: error while selecting actor info: %w", err)
			logger.Error().Err(errMsg).Msg("pg_error")

			return nil, errMsg
		}
	}

	metricsutils.SaveSuccessMetric(start, "update_subscription_status", "subscriptions")
	logger.Info().Msg("postgres: subscription status updated successfully")

	return sub, nil
}
