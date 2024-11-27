package service

// import (
// 	"context"

// 	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
// 	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
// )

// func (s *MovieService) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.ServiceError) {
// 	movies, err := s.movieClient.GetMovieByGenre(ctx, genre)

// 	if err != nil {
// 		return nil, errVals.NewServiceError("failed_to_reset_favorite", err)
// 	}

// 	return movies, nil
// }
