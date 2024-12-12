package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	movieDB "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/moviedb"
)

// GetFavorites gets favorites by calling db GetMoviesByIDs
func (r *MovieRepo) GetFavorites(ctx context.Context, mvIDs []uint64) ([]*models.MovieShortInfo, error) {
	var rows *sql.Rows
	var err error

	rows, err = movieDB.GetMoviesByIDs(ctx, mvIDs, r.Database)

	if err != nil {
		return nil, fmt.Errorf("error getting movie collections: %w", err)
	}

	favorites, err := movieDB.ScanMovieShortConnection(rows)
	if err != nil {
		return nil, fmt.Errorf("error getting movie collections: %w", err)
	}

	return favorites, nil
}
