package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/subscription/delivery"
)

// SubscriptionService is a facade service layer struct
type SubscriptionService struct {
	paymentClient client.PaymentClientInterface
	usrClient     client.UserClientInterface
}

// NewSubscriptionService returns an instance of SubscriptionServiceInterface
func NewSubscriptionService(pc client.PaymentClientInterface, uc client.UserClientInterface) delivery.SubscriptionServiceInterface {
	return &SubscriptionService{
		paymentClient: pc,
		usrClient:     uc,
	}
}
