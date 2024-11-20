package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	api "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type MovieRepositoryInterface interface {
	GetCollection(ctx context.Context) ([]models.Collection, *errVals.ErrorObj, int)
	GetMovie(ctx context.Context, mvId int) (*models.MovieInfo, *errVals.ErrorObj, int)
	GetActor(ctx context.Context, actorId int) (*models.StaffInfo, *errVals.ErrorObj, int)
	SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error)
	SearchActors(ctx context.Context, query string) ([]models.StaffInfo, error)
}

type MovieService struct {
	movieRepository MovieRepositoryInterface
}

func NewService(repo MovieRepositoryInterface) api.MovieServiceInterface {
	return &MovieService{
		movieRepository: repo,
	}
}
