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
		INSERT INTO subscriptions (user_id, price, status, expiration_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	findByUserIDSQL = `
		SELECT status, expiration_date
		FROM subscriptions
		WHERE user_id = $1 and expiration_date > $2 and status = $3
	`
	markPaidSQL = "UPDATE subscriptions SET status = $1, expiration_date = $2 WHERE id = $3"

	pendingStatus = "pending"
	// ActiveStatus defines subscription active state
	ActiveStatus = "active"
)

// CreateSubscription creates subscription in db
func CreateSubscription(ctx context.Context, subData *dto.RepoCreateSubscriptionData, db *sql.DB) (uint64, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	var subID uint64
	stmt, err := db.Prepare(subCreateSQL)
	if err != nil {
		return 0, fmt.Errorf("prepareStatement#createSubscription: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	err = stmt.QueryRowContext(
		ctx,
		subData.UserID, subData.Amount, pendingStatus, time.Now().AddDate(0, 1, 0),
	).Scan(&subID)

	if err != nil {
		metricsutils.SaveErrorMetric("create_subscription", "subscriptions")
		errMsg := fmt.Errorf("postgres: error while creating subscription - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return 0, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "create_subscription", "subscriptions")
	logger.Info().Msg("postgres: subscription created successfully")

	return subID, nil
}

// UpdateSubscription updates subscription data in db
func UpdateSubscription(ctx context.Context, subID uint64, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(markPaidSQL)
	if err != nil {
		return fmt.Errorf("prepareStatement#updateSubscription: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	_, err = stmt.ExecContext(ctx, ActiveStatus, time.Now().AddDate(0, 1, 0), subID)

	if err != nil {
		metricsutils.SaveErrorMetric("update_subscription_status", "subscriptions")
		errMsg := fmt.Errorf("postgres: error while updating subscription status - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "update_subscription_status", "subscriptions")
	logger.Info().Msg("postgres: subscription status updated successfully")

	return nil
}

// FindByUserID gets last active user subscription in db
func FindByUserID(ctx context.Context, usrID uint64, db *sql.DB) (*dto.RepoSubscription, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	var sub = &dto.RepoSubscription{}
	stmt, err := db.Prepare(findByUserIDSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#subscriptionByUserID: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	row := stmt.QueryRowContext(ctx, usrID, time.Now(), ActiveStatus)

	err = row.Scan(
		&sub.Status,
		&sub.ExpirationDate,
	)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			metricsutils.SaveErrorMetric("get_actor_by_id", "actors")
			errMsg := fmt.Errorf("postgres: error while selecting actor info: %w", err)
			logger.Error().Err(errMsg).Msg("pg_error")

			return nil, errMsg
		}
	}

	metricsutils.SaveSuccessMetric(start, "update_subscription_status", "subscriptions")
	logger.Info().Msg("postgres: subscription status updated successfully")

	return sub, nil
}
