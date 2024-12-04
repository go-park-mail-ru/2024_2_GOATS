package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/paymentdb"
)

func (u *PaymentRepo) CreatePayment(ctx context.Context, paymentData *dto.RepoPaymentData) (uint64, error) {
	pID, err := paymentdb.Create(ctx, paymentData, u.Database)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errors.ErrCreatePaymentCode, err)
	}

	return pID, nil
}
