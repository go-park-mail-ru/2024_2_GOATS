package moviedb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/metrics_utils"
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
			countries.title,
			episodes.id,
   		episodes.title,
   		episodes.description,
   		seasons.season_number,
   		episodes.episode_number,
   		episodes.release_date,
   		episodes.rating,
   		episodes.preview_url,
   		episodes.video_url
		FROM movies
		JOIN directors ON directors.id = movies.director_id
		JOIN countries ON countries.id = movies.country_id
		LEFT JOIN seasons ON seasons.movie_id = movies.id AND movies.movie_type = 'serial'
		LEFT JOIN episodes ON seasons.id = episodes.season_id AND movies.movie_type = 'serial'
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

	getFavoritesSQL = `
		SELECT movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM movies
		JOIN countries ON countries.id = movies.country_id
		WHERE movies.id = ANY($1)
	`
)

func FindByID(ctx context.Context, mvID int, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, movieFindByIDSQL, mvID)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "get_movie_by_id", "movies")
		errMsg := fmt.Errorf("postgres: error while selecting movie_service info: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_movie_by_id", "movies")
	logger.Info().Msg("postgres: successfully select movie_service info")

	return rows, nil
}

func GetMovieActors(ctx context.Context, mvID int, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, getMovieActorsSQL, mvID)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "get_movie_actors", "actors")
		errMsg := fmt.Errorf("postgres: error while selecting movie_service actors info: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_movie_actors", "actors")
	logger.Info().Msg("postgres: successfully select movie_service actors info")

	return rows, nil
}

func FindByActorID(ctx context.Context, actorID int, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, findByActorIDSQL, actorID)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "find_by_actor_id", "movies")
		errMsg := fmt.Errorf("postgres: error while selecting actor's movies: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "find_by_actor_id", "movies")
	return rows, nil
}

func GetMoviesByIDs(ctx context.Context, mvIDs []uint64, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	rows, err := db.Query(getFavoritesSQL, pq.Array(mvIDs))

	if err != nil {
		metricsutils.SaveErrorMetric(start, "get_movie_by_ids", "movies")
		errMsg := fmt.Errorf("postgres: error while selecting favorite movies: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_movie_by_ids", "movies")
	return rows, nil
}
