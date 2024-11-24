package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	api "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
	usrSrv "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type MovieRepositoryInterface interface {
	GetCollection(ctx context.Context, filter string) ([]models.Collection, *errVals.RepoError)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.RepoError)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.RepoError)
	GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, *errVals.RepoError)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.RepoError)
}

type MovieService struct {
	movieRepository MovieRepositoryInterface
	userRepository  usrSrv.UserRepositoryInterface
}

func NewMovieService(repo MovieRepositoryInterface, urepo usrSrv.UserRepositoryInterface) api.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
		userRepository:  urepo,
	}
}
