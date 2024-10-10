package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service"
	"github.com/go-redis/redis/v8"
)

var _ service.RepositoryInterface = (*Repo)(nil)

type Repo struct {
	Database *sql.DB
	Redis    *redis.Client
}

func NewRepository(db *sql.DB, rdb *redis.Client) service.RepositoryInterface {
	return &Repo{
		Database: db,
		Redis:    rdb,
	}
}
