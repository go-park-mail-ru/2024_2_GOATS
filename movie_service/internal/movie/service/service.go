package service

import (
	"context"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
)

// MovieRepositoryInterface defines methods for movie_service repo layer
//
//go:generate mockgen -source=service.go -destination=mocks/mock.go
type MovieRepositoryInterface interface {
	GetCollection(ctx context.Context, filter string) ([]models.Collection, error)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error)
	GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error)
	SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error)
	SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error)
	GetFavorites(ctx context.Context, mvIDs []uint64) ([]*models.MovieShortInfo, error)
}

// MovieService is a movie_service service layer struct
type MovieService struct {
	movieRepository MovieRepositoryInterface
}

// NewMovieService returns an instance of MovieServiceInterface
func NewMovieService(repo MovieRepositoryInterface) delivery.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
	}
}

// GetMovieByGenre gets movie by genre
func (s *MovieService) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	movies, err := s.movieRepository.GetMovieByGenre(ctx, genre)

	if err != nil {
		return nil, fmt.Errorf("movieService.GetMovieByGenre: %w", err)
	}

	return movies, nil
}

// GetCollection gets movie collections
func (s *MovieService) GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, error) {
	collections, err := s.movieRepository.GetCollection(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("movieService.GetCollection: %w", err)
	}

	return &models.CollectionsRespData{Collections: collections}, nil
}

// SearchMovies search movies via elastic
func (s *MovieService) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	return s.movieRepository.SearchMovies(ctx, query)
}

// SearchActors search actors via elastic
func (s *MovieService) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	return s.movieRepository.SearchActors(ctx, query)
}

// GetMovie gets movie by id
func (s *MovieService) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	mv, err := s.movieRepository.GetMovie(ctx, mvID)
	if err != nil {
		return nil, fmt.Errorf("movieService.GetMovie: %w", err)
	}

	actors, err := s.movieRepository.GetMovieActors(ctx, mv.ID)

	if err != nil {
		return nil, fmt.Errorf("movieService.GetMovieActors: %w", err)
	}

	mv.Actors = actors

	return mv, nil
}

// GetActor gets actor by id
func (s *MovieService) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	actor, err := s.movieRepository.GetActor(ctx, actorID)
	log.Println("actorServ", actor)
	if err != nil {
		return nil, fmt.Errorf("movieService.GetActor: %w", err)
	}

	return actor, nil
}

// GetMovieActors gets movie's actors by id
func (s *MovieService) GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error) {
	actor, err := s.movieRepository.GetMovieActors(ctx, mvID)

	if err != nil {
		return nil, fmt.Errorf("movieService.GetMovieActors: %w", err)
	}

	return actor, nil
}

// GetFavorites gets favorite movies by ids
func (s *MovieService) GetFavorites(ctx context.Context, mvIDs []uint64) ([]*models.MovieShortInfo, error) {
	mvs, err := s.movieRepository.GetFavorites(ctx, mvIDs)
	if err != nil {
		return nil, fmt.Errorf("movieService.GetFavorites: %w", err)
	}

	return mvs, nil
}
