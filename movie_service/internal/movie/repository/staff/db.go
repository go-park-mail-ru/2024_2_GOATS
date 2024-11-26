package staff

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
)

func FindById(ctx context.Context, staffId int, post string, db *sql.DB) *sql.Row {
	logger := log.Ctx(ctx)

	actorSqlStatement := `
		SELECT
			movie_staff.id,
			movie_staff.first_name,
			movie_staff.second_name,
			movie_staff.biography,
			movie_staff.birthdate,
			movie_staff.big_photo_url,
			countries.title
		FROM movie_staff
		JOIN countries on countries.id = movie_staff.country_id
		WHERE movie_staff.id = $1 and movie_staff.post = $2
	`

	row := db.QueryRowContext(ctx, actorSqlStatement, staffId, post)

	logger.Info().Msg("postgres: successfully select staff info")

	return row
}
