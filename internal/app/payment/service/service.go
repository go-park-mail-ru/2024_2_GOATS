package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/payment/delivery"
)

type PaymentService struct {
	paymentClient client.PaymentClientInterface
	usrClient     client.UserClientInterface
}

func NewPaymentService(pc client.PaymentClientInterface, uc client.UserClientInterface) delivery.PaymentServiceInterface {
	return &PaymentService{
		paymentClient: pc,
		usrClient:     uc,
	}
}
