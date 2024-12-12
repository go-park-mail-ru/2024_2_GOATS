package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	movieCollectionDB "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/movie_collectiondb"
)

// GetCollection gets movie collections by calling db GetCollections
func (r *MovieRepo) GetCollection(ctx context.Context, filter string) ([]models.Collection, error) {
	var rows *sql.Rows
	var err error

	if filter == "genres" {
		rows, err = movieCollectionDB.GetGenreCollections(ctx, r.Database)
	} else {
		rows, err = movieCollectionDB.GetMovieCollections(ctx, r.Database)
	}

	if err != nil {
		return nil, fmt.Errorf("error getting movie collections: %w", err)
	}

	collections, err := movieCollectionDB.ScanConnections(rows)
	if err != nil {
		return nil, fmt.Errorf("error getting movie collections: %w", err)
	}

	result := make([]models.Collection, 0, len(collections))
	for _, collection := range collections {
		result = append(result, collection)
	}

	return result, nil
}
