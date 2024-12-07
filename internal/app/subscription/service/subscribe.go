package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func (ss *SubscriptionService) Subscribe(ctx context.Context, data *models.SubscriptionData) (string, *errVals.ServiceError) {
	logger := log.Ctx(ctx)
	subID, err := ss.usrClient.CreateSubscription(ctx, data)
	if err != nil {
		errMsg := fmt.Errorf("create_subscription_error: %w", err)
		logger.Error().Err(errMsg).Msg("failed_to_create_subscription")

		return "", errVals.NewServiceError("failed_to_create_subscription", errMsg)
	}

	pd := &models.CreatePaymentData{
		SubscriptionID: subID,
		Amount:         data.Amount,
	}

	pID, err := ss.paymentClient.CreatePayment(ctx, pd)
	if err != nil {
		logger.Error().Err(err).Msg("failed_to_create_payment")

		return "", errVals.NewServiceError("failed_to_create_payment", err)
	}

	return fmt.Sprintf("%d-%d", subID, pID), nil
}
