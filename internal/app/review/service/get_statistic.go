package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/service/dto"
	"github.com/rs/zerolog/log"
)

func (rs *ReviewService) GetStatistics(ctx context.Context) (*dto.Statistic, *errors.ServiceError) {
	logger := log.Ctx(ctx)
	resp, err := rs.reviewClient.GetStatistics(ctx)

	if err != nil {
		errMsg := fmt.Errorf("cannot get questions: %w", err)
		logger.Error().Err(errMsg).Msg("get_questions_error")
		return nil, errors.NewServiceError("get_questions_error", errors.CustomError{Err: errMsg})
	}

	return resp, nil
}
