package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, *errVals.ServiceError) {
	collections, err := s.movieRepository.GetCollection(ctx, filter)

	if err != nil {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	return &models.CollectionsRespData{Collections: collections}, nil
}
