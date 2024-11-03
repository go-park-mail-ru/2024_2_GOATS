package movie_collection

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
)

func Obtain(ctx context.Context, db *sql.DB) (*sql.Rows, error) {
	logger, requestId := config.FromBaseContext(ctx)

	sqlStatement := `
	SELECT collections.id, collections.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM collections
	JOIN movie_collections ON movie_collections.collection_id = collections.id
	JOIN movies ON movies.id = movie_collections.movie_id
	JOIN countries ON countries.id = movies.country_id
`

	rows, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting movie_collections: %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return nil, errMsg
	}

	logger.Log("postgres: successfully select movie_collections", requestId)

	return rows, nil
}
