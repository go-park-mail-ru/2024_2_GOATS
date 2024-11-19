package movie_collectiondb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	getMovieCollectionsSQL = `
		SELECT collections.id, collections.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM collections
		JOIN movie_collections ON movie_collections.collection_id = collections.id
		JOIN movies ON movies.id = movie_collections.movie_id
		JOIN countries ON countries.id = movies.country_id
	`

	getGenresCollectionsSQL = `
		SELECT genres.id, genres.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM genres
		JOIN movie_genres ON movie_genres.genre_id = genres.id
		JOIN movies ON movies.id = movie_genres.movie_id
		JOIN countries ON countries.id = movies.country_id
	`

	getByGenreSQL = `
		SELECT movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM movies
		JOIN movie_genres ON movie_genres.movie_id = movies.id
		JOIN genres ON genres.id = movie_genres.genre_id AND genres.title = $1
		JOIN countries ON countries.id = movies.country_id
	`
)

func GetMovieCollections(ctx context.Context, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, getMovieCollectionsSQL)
	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie_collections: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie_collections")

	return rows, nil
}

func GetGenreCollections(ctx context.Context, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, getGenresCollectionsSQL)
	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting genre_collections: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie_collections")

	return rows, nil
}

func GetMovieByGenre(ctx context.Context, genre string, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, getByGenreSQL, genre)
	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movies by genre: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movies by genre")

	return rows, nil
}
