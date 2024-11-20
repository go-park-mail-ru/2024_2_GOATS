package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	return s.movieRepository.SearchMovies(ctx, query)
}

func (s *MovieService) SearchActors(ctx context.Context, query string) ([]models.StaffInfo, error) {
	return s.movieRepository.SearchActors(ctx, query)
}
