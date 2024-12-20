package moviedb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	"time"

	metricsutils "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/metrics_utils"
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
			movies.with_subscription,
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

	getGenresSQL = `
		SELECT genres.title FROM genres
		JOIN movie_genres on movie_genres.genre_id = genres.id
		JOIN movies ON movies.id = movie_genres.movie_id
		WHERE movies.id = $1
	`
)

// FindByID finds movie by id in db
func FindByID(ctx context.Context, mvID int, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(movieFindByIDSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#movieByID: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	rows, err := stmt.QueryContext(ctx, mvID)

	if err != nil {
		metricsutils.SaveErrorMetric("get_movie_by_id", "movies")
		errMsg := fmt.Errorf("postgres: error while selecting movie_service info: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_movie_by_id", "movies")
	logger.Info().Msg("postgres: successfully select movie_service info")

	return rows, nil
}

// GetMovieActors finds movie's actors in db
func GetMovieActors(ctx context.Context, mvID int, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(getMovieActorsSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#actorsByMovieID: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	rows, err := stmt.QueryContext(ctx, mvID)

	if err != nil {
		metricsutils.SaveErrorMetric("get_movie_actors", "actors")
		errMsg := fmt.Errorf("postgres: error while selecting movie_service actors info: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_movie_actors", "actors")
	logger.Info().Msg("postgres: successfully select movie_service actors info")

	return rows, nil
}

// FindByActorID finds movies by actor_id in db
func FindByActorID(ctx context.Context, actorID int, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(findByActorIDSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#moviesByActorID: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	rows, err := stmt.QueryContext(ctx, actorID)

	if err != nil {
		metricsutils.SaveErrorMetric("find_by_actor_id", "movies")
		errMsg := fmt.Errorf("postgres: error while selecting actor's movies: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "find_by_actor_id", "movies")
	return rows, nil
}

// GetMoviesByIDs finds movies by ids in db
func GetMoviesByIDs(ctx context.Context, mvIDs []uint64, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(getFavoritesSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#moviesByIDs: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	rows, err := stmt.QueryContext(ctx, pq.Array(mvIDs))

	if err != nil {
		metricsutils.SaveErrorMetric("get_movie_by_ids", "movies")
		errMsg := fmt.Errorf("postgres: error while selecting favorite movies: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_movie_by_ids", "movies")
	return rows, nil
}

// GetGenres returns genres for film
func GetGenres(ctx context.Context, mvID int, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(getGenresSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#moviesByIDs: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	rows, err := stmt.QueryContext(ctx, mvID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		metricsutils.SaveErrorMetric("get_movie_genres", "genres")
		errMsg := fmt.Errorf("postgres: error while selecting movie genres: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_movie_genres", "movies")
	return rows, nil
}
