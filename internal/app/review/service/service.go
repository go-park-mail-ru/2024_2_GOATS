package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/delivery"
)

type ReviewService struct {
	reviewClient client.ReviewClientInterface
}

func NewReviewService(client client.ReviewClientInterface) delivery.ReviewServiceInterface {
	return &ReviewService{
		reviewClient: client,
	}
}
