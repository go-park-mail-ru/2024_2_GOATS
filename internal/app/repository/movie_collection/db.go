package movie_collection

import (
	"context"
	"database/sql"
	"fmt"
)

func Obtain(ctx context.Context, db *sql.DB) (*sql.Rows, error) {
	sqlStatement := `
	SELECT collections.id, collections.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, countries.title FROM collections
	JOIN movie_collections ON movie_collections.collection_id = collections.id
	JOIN movies ON movies.id = movie_collections.movie_id
	JOIN countries ON countries.id = movies.country_id
`

	rows, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error while scanning movie_collections: %w", err)
	}

	return rows, nil
}
