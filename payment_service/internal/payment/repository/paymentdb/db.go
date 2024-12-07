package paymentdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/dto"
	metricsutils "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/metrics_utils"
	"github.com/rs/zerolog/log"
)

const (
	usrCreateSQL = `
		INSERT INTO payments (subscription_id, requested_amount)
		VALUES ($1, $2)
		RETURNING id
	`

	markPaidSQL = "UPDATE payments SET captured_amount = requested_amount, captured_at = $1 WHERE id = $2"
)

func Create(ctx context.Context, paymentData *dto.RepoPaymentData, db *sql.DB) (uint64, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	var pID uint64
	err := db.QueryRowContext(
		ctx,
		usrCreateSQL,
		paymentData.SubscriptionID, paymentData.Amount,
	).Scan(&pID)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "create_payment", "payments")
		errMsg := fmt.Errorf("postgres: error while creating payment - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return 0, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "create_payment", "payments")
	logger.Info().Msg("postgres: payment created successfully")

	return pID, nil
}

func MarkPaid(ctx context.Context, pID uint64, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	_, err := db.ExecContext(
		ctx,
		markPaidSQL,
		time.Now(), pID,
	)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "mark_paid_payment", "payments")
		errMsg := fmt.Errorf("postgres: error while marking payment as paid - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "mark_paid_payment", "payments")
	logger.Info().Msg("postgres: payment marked paid successfully")

	return nil
}
