package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service"
	"github.com/go-redis/redis/v8"
)

var _ service.MovieRepositoryInterface = (*Repo)(nil)

type Repo struct {
	Database *sql.DB
	Redis    *redis.Client
}

func NewRepository(db *sql.DB, rdb *redis.Client) service.MovieRepositoryInterface {
	return &Repo{
		Database: db,
		Redis:    rdb,
	}
}
