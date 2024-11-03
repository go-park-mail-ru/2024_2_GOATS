package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service"
	"github.com/go-redis/redis/v8"
)

var _ service.MovieRepositoryInterface = (*MovieRepo)(nil)

type MovieRepo struct {
	Database *sql.DB
	Redis    *redis.Client
}

func NewMovieRepository(db *sql.DB) service.MovieRepositoryInterface {
	return &MovieRepo{
		Database: db,
	}
}
