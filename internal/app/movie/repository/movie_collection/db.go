package movie_collection

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

func Obtain(ctx context.Context, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	sqlStatement := `
	SELECT collections.id, collections.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM collections
	JOIN movie_collections ON movie_collections.collection_id = collections.id
	JOIN movies ON movies.id = movie_collections.movie_id
	JOIN countries ON countries.id = movies.country_id
`

	rows, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie_collections: %w", err)
		logger.Error().Msg(errMsg.Error())

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select movie_collections")

	return rows, nil
}
