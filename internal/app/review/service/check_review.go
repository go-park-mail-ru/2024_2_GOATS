package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/service/dto"
	"github.com/rs/zerolog/log"
)

func (rs *ReviewService) CheckReview(ctx context.Context, usrID int) (*dto.CheckReviewData, *errors.ServiceError) {
	logger := log.Ctx(ctx)
	pr, err := rs.reviewClient.CheckReview(ctx, usrID)
	if err != nil {
		errMsg := fmt.Errorf("cannot check review existence: %w", err)
		logger.Error().Err(errMsg).Msg("check_review_error")
		return nil, errors.NewServiceError("check_review_error", errors.CustomError{Err: errMsg})
	}
	return &dto.CheckReviewData{
		CSAT: pr,
	}, nil
}
