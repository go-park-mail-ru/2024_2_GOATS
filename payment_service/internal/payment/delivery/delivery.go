package delivery

import (
	"context"

	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/dto"
)

// PaymentServiceInterface defines payment_service methods
//
//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type PaymentServiceInterface interface {
	CreatePayment(ctx context.Context, createData *srvDTO.CreatePaymentData) (uint64, error)
	MarkPaid(ctx context.Context, pID uint64) error
}
