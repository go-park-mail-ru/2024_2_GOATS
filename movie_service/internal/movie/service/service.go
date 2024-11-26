package service

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	usrSrv "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/clients"
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
}

type Favorite struct {
	UserID  uint64
	MovieID uint64
}

type MovieService struct {
	movieRepository MovieRepositoryInterface
	userClient      usrSrv.UserClientInterface
}

func NewMovieService(repo MovieRepositoryInterface, urepo usrSrv.UserClientInterface) delivery.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
		userClient:      urepo,
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

	usrID := config.CurrentUserID(ctx)
	if usrID != 0 {
		fav := &models.Favorite{
			UserID:  usrID,
			MovieID: mv.ID,
		}

		isFav, err := s.userClient.CheckFavorite(ctx, fav)
		if err != nil {
			return nil, fmt.Errorf("movieService.GetMovie: %w", err)
		}

		mv.IsFavorite = isFav
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
