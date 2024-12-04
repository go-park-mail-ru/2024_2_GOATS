package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/dto"
	"github.com/rs/zerolog/log"
)

func (u *PaymentService) CreatePayment(ctx context.Context, createData *dto.CreatePaymentData) (uint64, error) {
	logger := log.Ctx(ctx)
	repoData := converter.ConvertToRepoPaymentData(createData)
	pID, err := u.paymentRepo.CreatePayment(ctx, repoData)
	if err != nil {
		logger.Error().Err(err).Msg("paymentService - failed to create payment")
		return 0, fmt.Errorf("paymentService - failed to create payment: %w", err)
	}

	return pID, nil
}
