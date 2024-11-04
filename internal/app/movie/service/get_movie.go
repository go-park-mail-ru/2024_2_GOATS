package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *models.ErrorRespData) {
	mv, err, code := s.movieRepository.GetMovie(ctx, mvID)

	if err != nil {
		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     []errVals.ErrorObj{*err},
		}
	}

	actors, err, code := s.movieRepository.GetMovieActors(ctx, mv.ID)

	if err != nil {
		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     []errVals.ErrorObj{*err},
		}
	}

	mv.Actors = actors

	return mv, nil
}
