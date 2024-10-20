package movie

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func FindById(ctx context.Context, mvId int, db *sql.DB) (*models.MovieFullData, *sql.Rows, error) {
	logger := log.Ctx(ctx)
	mv := &models.MovieFullData{}
	mvInfo := &models.MovieBaseInfo{}

	mvSqlStatement := `
		SELECT
			movies.id,
			movies.title,
			movies.description,
			movies.card_url,
			movies.album_url,
			movies.rating,
			movies.release_date,
			movies.video_url,
			movies.movie_type,
			countries.title
		FROM movies
		JOIN countries ON countries.id = movies.country_id
		WHERE movies.id = $1
	`

	err := db.QueryRowContext(ctx, mvSqlStatement, mvId).
		Scan(
			&mvInfo.Id,
			&mvInfo.Title,
			&mvInfo.Description,
			&mvInfo.CardUrl,
			&mvInfo.AlbumUrl,
			&mvInfo.Rating,
			&mvInfo.ReleaseDate,
			&mvInfo.VideoUrl,
			&mvInfo.MovieType,
			&mvInfo.Country,
		)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie info: %w", err)
		logger.Err(errMsg)

		return nil, nil, errMsg
	}

	mv.MovieBaseInfo = mvInfo
	logger.Info().Msg("postgres: successfully select movie info")

	actorsSqlStatement := `
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.patronymic,
			actors.photo_url
		FROM actors
		JOIN movie_actors on movie_actors.actor_id = actors.id
		WHERE movie_actors.movie_id = $1
	`

	rows, err := db.QueryContext(ctx, actorsSqlStatement, mvId)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie info: %w", err)
		logger.Err(errMsg)

		return mv, nil, errMsg
	}

	return mv, rows, nil
}
