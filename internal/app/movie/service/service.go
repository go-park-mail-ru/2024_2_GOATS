package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	api "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
)

var _ api.MovieServiceInterface = (*MovieService)(nil)

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
}

func NewMovieService(repo MovieRepositoryInterface) api.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
	}
}
