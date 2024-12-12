package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/paymentdb"
)

// MarkPaid calls db MarkPaid
func (u *PaymentRepo) MarkPaid(ctx context.Context, pID uint64) error {
	err := paymentdb.MarkPaid(ctx, pID, u.Database)
	if err != nil {
		return fmt.Errorf("%s: %w", errors.ErrMarkPaidPaymentCode, err)
	}

	return nil
}
