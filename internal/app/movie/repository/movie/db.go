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
			movie_staff.id,
			movie_staff.first_name,
			movie_staff.second_name,
			movie_staff.patronymic,
			movie_staff.biography,
			movie_staff.post,
			movie_staff.small_photo_url,
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
			countries.title
		FROM movies
		JOIN countries ON countries.id = movies.country_id
		join staff_members on staff_members.movie_id = movies.id
		join movie_staff on staff_members.movie_staff_id = movie_staff.id
		WHERE movies.id = $1
	`

	rows, err := db.QueryContext(ctx, mvSqlStatement, mvId)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie info: %w", err)
		logger.Err(errMsg)

		return nil, errMsg
	}

	hasRows := rows.Next()
	if !hasRows {
		errMsg := fmt.Errorf("postgres: error while select movie info: %w", sql.ErrNoRows)
		logger.Err(errMsg)

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie info")

	return rows, nil
}
