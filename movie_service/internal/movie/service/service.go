package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"

	usrSrv "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	api "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type MovieRepositoryInterface interface {
	GetCollection(ctx context.Context, filter string) ([]models.Collection, *errVals.RepoError)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.RepoError)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.RepoError)
	GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, *errVals.RepoError)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.RepoError)
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

func NewMovieService(repo MovieRepositoryInterface, urepo usrSrv.UserClientInterface) api.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
		userClient:      urepo,
	}
}

func (s *MovieService) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	movies, err := s.movieRepository.GetMovieByGenre(ctx, genre)

	if err != nil {
		return nil, &errVals.ServiceError{Code: "1"}
	}

	return movies, nil
}

func (s *MovieService) GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, error) {
	collections, err := s.movieRepository.GetCollection(ctx, filter)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code: "1",
		}
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
		return nil, &errVals.ServiceError{
			Code: "1",
		}
	}

	usrID := config.CurrentUserID(ctx)
	if usrID != 0 {
		fav := &models.Favorite{
			UserID:  usrID,
			MovieID: mv.ID,
		}

		isFav, err := s.userClient.CheckFavorite(ctx, fav)
		if err != nil {
			return nil, &errVals.ServiceError{
				Code: "1",
			}
		}

		mv.IsFavorite = isFav
	}

	actors, err := s.movieRepository.GetMovieActors(ctx, mv.ID)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code: "1",
		}
	}

	mv.Actors = actors

	return mv, nil
}

func (s *MovieService) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	actor, err := s.movieRepository.GetActor(ctx, actorID)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code: "1",
		}
	}

	return actor, nil
}

func (s *MovieService) GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error) {
	actor, err := s.movieRepository.GetMovieActors(ctx, mvID)

	if err != nil {
		return nil, err.Error.Err
	}

	return actor, nil
}
