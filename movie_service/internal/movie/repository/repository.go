package repository

import (
	"database/sql"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/service"
)

type MovieRepo struct {
	Database      *sql.DB
	Elasticsearch *elasticsearch.Client
}

func NewMovieRepository(db *sql.DB, es *elasticsearch.Client) service.MovieRepositoryInterface {
	return &MovieRepo{
		Database:      db,
		Elasticsearch: es,
	}
}
