package service

import (
	"context"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
)

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
	GetUserRating(ctx context.Context, userId int, movieId int) (float32, error)
	AddOrUpdateRating(ctx context.Context, userId int, movieId int, rating float32) error
	DeleteUserRating(ctx context.Context, userID, movieID int) error
	UpdateMovieRating(ctx context.Context, movieId int) error
}

type Favorite struct {
	UserID  uint64
	MovieID uint64
}

type MovieService struct {
	movieRepository MovieRepositoryInterface
}

func NewMovieService(repo MovieRepositoryInterface) delivery.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
	}
}

func (s *MovieService) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	movies, err := s.movieRepository.GetMovieByGenre(ctx, genre)

	if err != nil {
		return nil, fmt.Errorf("movieService.GetMovieByGenre: %w", err)
	}

	return movies, nil
}

func (s *MovieService) GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, error) {
	collections, err := s.movieRepository.GetCollection(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("movieService.GetCollection: %w", err)
	}

	return &models.CollectionsRespData{Collections: collections}, nil
}

func (s *MovieService) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	return s.movieRepository.SearchMovies(ctx, query)
}

func (s *MovieService) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	return s.movieRepository.SearchActors(ctx, query)
}

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

func (s *MovieService) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	actor, err := s.movieRepository.GetActor(ctx, actorID)
	log.Println("actorServ", actor)
	if err != nil {
		return nil, fmt.Errorf("movieService.GetActor: %w", err)
	}

	return actor, nil
}

func (s *MovieService) GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error) {
	actor, err := s.movieRepository.GetMovieActors(ctx, mvID)

	if err != nil {
		return nil, fmt.Errorf("movieService.GetMovieActors: %w", err)
	}

	return actor, nil
}

func (s *MovieService) GetFavorites(ctx context.Context, mvIDs []uint64) ([]*models.MovieShortInfo, error) {
	mvs, err := s.movieRepository.GetFavorites(ctx, mvIDs)
	if err != nil {
		return nil, fmt.Errorf("movieService.GetFavorites: %w", err)
	}

	return mvs, nil
}

func (s *MovieService) GetUserRating(ctx context.Context, userId int, movieId int) (float32, error) {
	rating, err := s.movieRepository.GetUserRating(ctx, userId, movieId)
	if err != nil {
		return 0, fmt.Errorf("movieService.GetUserRating: %w", err)
	}

	log.Println("rating", rating)
	return rating, nil
}

func (s *MovieService) AddOrUpdateRating(ctx context.Context, userId int, movieId int, rating float32) error {
	err := s.movieRepository.AddOrUpdateRating(ctx, userId, movieId, rating)
	if err != nil {
		return fmt.Errorf("movieService.AddOrUpdateRating: %w", err)
	}

	err = s.movieRepository.UpdateMovieRating(ctx, movieId)
	if err != nil {
		return fmt.Errorf("movieService.AddOrUpdateRating (update movie rating): %w", err)
	}

	return nil
}

func (s *MovieService) DeleteRating(ctx context.Context, userID, movieID int) error {
	err := s.movieRepository.DeleteUserRating(ctx, userID, movieID)
	if err != nil {
		return fmt.Errorf("movieService.DeleteRating: %w", err)
	}
	return nil
}
