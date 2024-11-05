package moviedb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	movieFindByIDSQL = `
		SELECT
			movies.id,
			movies.title,
			movies.short_description,
			movies.long_description,
			movies.card_url,
			movies.album_url,
			movies.rating,
			movies.release_date,
			movies.video_url,
			movies.movie_type,
			movies.title_url,
			directors.first_name,
			directors.second_name,
			countries.title
		FROM movies
		JOIN directors ON directors.id = movies.director_id
		JOIN countries ON countries.id = movies.country_id
		WHERE movies.id = $1
	`

	getMovieActorsSQL = `
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.biography,
			actors.small_photo_url
		FROM actors
		JOIN movie_actors on movie_actors.actor_id = actors.id
		JOIN movies on movie_actors.movie_id = movies.id
		WHERE movies.id = $1
	`

	findByActorIDSQL = `
		SELECT
			movies.id,
			movies.title,
			movies.card_url,
			movies.rating,
			movies.release_date,
			countries.title
		FROM movies
		JOIN movie_actors ON movie_actors.movie_id = movies.id
		JOIN actors ON movie_actors.actor_id = actors.id
		JOIN countries ON movies.country_id = countries.id
		WHERE actors.id = $1
	`
)

func FindByID(ctx context.Context, mvID int, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, movieFindByIDSQL, mvID)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie info: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie info")

	return rows, nil
}

func GetMovieActors(ctx context.Context, mvID int, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, getMovieActorsSQL, mvID)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie actors info: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie actors info")

	return rows, nil
}

func FindByActorID(ctx context.Context, actorID int, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, findByActorIDSQL, actorID)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting actor's movies: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	return rows, nil
}
