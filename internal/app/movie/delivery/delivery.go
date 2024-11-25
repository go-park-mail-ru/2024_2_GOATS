package delivery

import (
	"context"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type MovieServiceInterface interface {
	GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, *errVals.ServiceError)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ServiceError)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.ServiceError)
	SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error)
	SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error)
}
