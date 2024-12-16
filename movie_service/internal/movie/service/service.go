package service

import (
	"context"
	"fmt"

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
	GetUserRating(ctx context.Context, userID int, movieID int) (float32, error)
	AddOrUpdateRating(ctx context.Context, userID int, movieID int, rating float32) error
	DeleteUserRating(ctx context.Context, userID, movieID int) error
	UpdateMovieRating(ctx context.Context, movieID int) error
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

// GetUserRating получение рейтинга пользователя
func (s *MovieService) GetUserRating(ctx context.Context, userID int, movieID int) (float32, error) {
	rating, err := s.movieRepository.GetUserRating(ctx, userID, movieID)
	if err != nil {
		return 0, fmt.Errorf("movieService.GetUserRating: %w", err)
	}

	return rating, nil
}

// AddOrUpdateRating добавление или обновление рейтинга
func (s *MovieService) AddOrUpdateRating(ctx context.Context, userID int, movieID int, rating float32) error {
	err := s.movieRepository.AddOrUpdateRating(ctx, userID, movieID, rating)
	if err != nil {
		return fmt.Errorf("movieService.AddOrUpdateRating: %w", err)
	}

	err = s.movieRepository.UpdateMovieRating(ctx, movieID)
	if err != nil {
		return fmt.Errorf("movieService.AddOrUpdateRating (update movie rating): %w", err)
	}

	return nil
}

// DeleteRating удаление рейтинга
func (s *MovieService) DeleteRating(ctx context.Context, userID, movieID int) error {
	err := s.movieRepository.DeleteUserRating(ctx, userID, movieID)
	if err != nil {
		return fmt.Errorf("movieService.DeleteRating: %w", err)
	}
	return nil
}
