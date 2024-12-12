package staff

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
)

// FindByID finds staff by id
func FindByID(ctx context.Context, staffID int, post string, db *sql.DB) *sql.Row {
	logger := log.Ctx(ctx)

	actorSQLStatement := `
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

	stmt, err := db.Prepare(actorSQLStatement)
	if err != nil {
		return nil
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	row := stmt.QueryRowContext(ctx, staffID, post)

	logger.Info().Msg("postgres: successfully select staff info")

	return row
}
