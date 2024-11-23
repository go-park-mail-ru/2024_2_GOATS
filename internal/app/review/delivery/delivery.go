package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/service/dto"
)

type ReviewServiceInterface interface {
	GetQuestions(ctx context.Context) ([]*dto.ReviewData, *errors.ServiceError)
	CheckReview(ctx context.Context, usrID int) (*dto.CheckReviewData, *errors.ServiceError)
	CreateCSAT(ctx context.Context, csatData []*dto.CreateReviewData) *errors.ServiceError
	GetStatistics(ctx context.Context) (*dto.Statistic, *errors.ServiceError)
}
