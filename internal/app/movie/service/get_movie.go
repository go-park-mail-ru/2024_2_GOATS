package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetMovie(ctx context.Context, mvId int) (*models.MovieInfo, *models.ErrorRespData) {
	mv, err, code := s.movieRepository.GetMovie(ctx, mvId)

	if err != nil {
		errs := make([]errVals.ErrorObj, 1)
		errs[0] = *err

		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     errs,
		}
	}

	return mv, nil
}
