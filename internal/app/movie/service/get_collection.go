package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetCollection(ctx context.Context) (*models.CollectionsRespData, *models.ErrorRespData) {
	collections, err, code := s.movieRepository.GetCollection(ctx)

	if err != nil {
		errs := make([]errVals.ErrorObj, 1)
		errs[0] = *err

		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     errs,
		}
	}

	return &models.CollectionsRespData{
		Collections: collections,
		StatusCode:  code,
	}, nil
}
