package converter

import (
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/dto"
	payment "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/pkg/payment_v1"
)

func ConvertToSrvPayment(req *payment.CreateRequest) *srvDTO.CreatePaymentData {
	if req == nil {
		return nil
	}

	return &srvDTO.CreatePaymentData{
		SubscriptionID: req.SubscriptionID,
		Amount:         req.Amount,
	}
}
