package repository

import (
	"database/sql"
	"github.com/elastic/go-elasticsearch/v7"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service"
	"github.com/go-redis/redis/v8"
)

type Repo struct {
	Database      *sql.DB
	Redis         *redis.Client
	Elasticsearch *elasticsearch.Client
}

func NewRepository(db *sql.DB, rdb *redis.Client, es *elasticsearch.Client) service.MovieRepositoryInterface {
	return &Repo{
		Database:      db,
		Redis:         rdb,
		Elasticsearch: es,
	}
}
