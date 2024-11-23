package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/service/dto"
	"github.com/rs/zerolog/log"
)

func (rs *ReviewService) CreateCSAT(ctx context.Context, csatData []*dto.CreateReviewData) *errors.ServiceError {
	logger := log.Ctx(ctx)
	usrID := config.CurrentUserID(ctx)
	err := rs.reviewClient.Create(ctx, usrID, csatData)
	if err != nil {
		errMsg := fmt.Errorf("cannot create csat: %w", err)
		logger.Error().Err(errMsg).Msg("create_csat_error")
		return errors.NewServiceError("create_csat_error", errors.CustomError{Err: errMsg})
	}

	logger.Info().Msg("successfully create csat")
	return nil
}
