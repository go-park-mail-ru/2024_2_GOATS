package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.ServiceError) {
	movies, err := s.movieRepository.GetByGenre(ctx, genre)

	if err != nil {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	return movies, nil
}
