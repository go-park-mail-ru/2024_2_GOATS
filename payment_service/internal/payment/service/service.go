package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/delivery"
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/dto"
)

// PaymentRepoInterface define methods for payment_service repo layer
//
//go:generate mockgen -source=service.go -destination=mocks/mock.go
type PaymentRepoInterface interface {
	CreatePayment(ctx context.Context, paymentData *repoDTO.RepoPaymentData) (uint64, error)
	MarkPaid(ctx context.Context, pID uint64) error
}

// PaymentService service layer struct
type PaymentService struct {
	paymentRepo PaymentRepoInterface
}

// NewPaymentService returns an instance of PaymentServiceInterface
func NewPaymentService(paymentRepo PaymentRepoInterface) delivery.PaymentServiceInterface {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}
