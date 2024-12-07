package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (u *PaymentService) MarkPaid(ctx context.Context, pID uint64) error {
	logger := log.Ctx(ctx)
	err := u.paymentRepo.MarkPaid(ctx, pID)
	if err != nil {
		logger.Error().Err(err).Msg("paymentService - failed to mark payment as paid")
		return fmt.Errorf("paymentService - failed to mark payment as paid: %w", err)
	}

	return nil
}
