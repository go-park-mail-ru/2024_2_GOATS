package delivery

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
)

// MovieServiceInterface defines methods for movie_service service layer
//
//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type MovieServiceInterface interface {
	GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, error)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error)
	SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error)
	SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error)
	GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error)
	GetFavorites(ctx context.Context, mvIDs []uint64) ([]*models.MovieShortInfo, error)
	GetUserRating(ctx context.Context, userId int, movieId int) (float32, error)
	AddOrUpdateRating(ctx context.Context, userId int, movieId int, rating float32) error
}
