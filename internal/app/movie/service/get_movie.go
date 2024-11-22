package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError) {
	// mv, err := s.movieRepository.GetMovie(ctx, mvID)

	// if err != nil {
	// 	return nil, errVals.ToServiceErrorFromRepo(err)
	// }

	// actors, err := s.movieRepository.GetMovieActors(ctx, mv.ID)

	// if err != nil {
	// 	return nil, errVals.ToServiceErrorFromRepo(err)
	// }

	// mv.Actors = actors

	// return mv, nil
	return nil, nil
}
