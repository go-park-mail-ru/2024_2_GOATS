package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"log"
)

func (s *MovieService) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	resp, err := s.movieRepository.SearchMovies(ctx, query)
	log.Println("movies", resp)
	log.Println("errService", err)
	return resp, err
}
