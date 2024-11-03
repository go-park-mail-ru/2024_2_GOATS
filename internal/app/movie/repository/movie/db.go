package movie

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

func FindById(ctx context.Context, mvId int, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	mvSqlStatement := `
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

	rows, err := db.QueryContext(ctx, mvSqlStatement, mvId)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie info: %w", err)
		logger.Error().Msg(errMsg.Error())

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie info")

	return rows, nil
}

func GetMovieActors(ctx context.Context, mvId int, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	actorsStatement := `
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

	rows, err := db.QueryContext(ctx, actorsStatement, mvId)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie actors info: %w", err)
		logger.Error().Msg(errMsg.Error())

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie actors info")

	return rows, nil
}

func FindByActorId(ctx context.Context, actorId int, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	mvStatement := `
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

	rows, err := db.QueryContext(ctx, mvStatement, actorId)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting actor's movies: %w", err)
		logger.Err(errMsg)

		return nil, errMsg
	}

	return rows, nil
}
