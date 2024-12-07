package converter

import (
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/dto"
)

func ConvertToRepoPaymentData(cr *dto.CreatePaymentData) *repoDTO.RepoPaymentData {
	if cr == nil {
		return nil
	}

	return &repoDTO.RepoPaymentData{
		SubscriptionID: cr.SubscriptionID,
		Amount:         cr.Amount,
	}
}
