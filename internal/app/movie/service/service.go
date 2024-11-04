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
	GetCollection(ctx context.Context) ([]models.Collection, *errVals.ErrorObj, int)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ErrorObj, int)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ErrorObj, int)
	GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, *errVals.ErrorObj, int)
}

type MovieService struct {
	movieRepository MovieRepositoryInterface
}

func NewMovieService(repo MovieRepositoryInterface) api.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
	}
}
