package service

import (
	"context"
	"strconv"
	"strings"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

// ProcessCallback processes payment callback by calling payment and user clients
func (ps *PaymentService) ProcessCallback(ctx context.Context, data *models.PaymentCallbackData) *errVals.ServiceError {
	logger := log.Ctx(ctx)

	labelSlice := strings.Split(data.Label, "-")
	strSubID, strPayID := labelSlice[0], labelSlice[1]
	payID, err := strconv.Atoi(strPayID)
	if err != nil {
		logger.Error().Err(err).Msg("failed_to_convert_callback_label")
		return errVals.NewServiceError("failed_to_convert_callback_label", err)
	}

	subID, err := strconv.Atoi(strSubID)
	if err != nil {
		logger.Error().Err(err).Msg("failed_to_convert_callback_label")
		return errVals.NewServiceError("failed_to_convert_callback_label", err)
	}

	err = ps.paymentClient.MarkPaid(ctx, payID)
	if err != nil {
		logger.Error().Err(err).Msg("failed_to_mark_payment_as_paid")

		return errVals.NewServiceError("failed_to_process_callback", err)
	}

	err = ps.usrClient.UpdateSubscriptionStatus(ctx, subID)
	if err != nil {
		logger.Error().Err(err).Msg("failed_to_update_subscription_status")

		return errVals.NewServiceError("failed_to_process_callback", err)
	}

	return nil
}
