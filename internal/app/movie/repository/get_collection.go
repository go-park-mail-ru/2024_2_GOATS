package repository

import (
	"context"
	"database/sql"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movieCollectionDB "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/movie_collectiondb"
)

func (r *MovieRepo) GetCollection(ctx context.Context, filter string) ([]models.Collection, *errVals.RepoError) {
	var rows *sql.Rows
	var err error

	if filter == "genres" {
		rows, err = movieCollectionDB.GetGenreCollections(ctx, r.Database)
	} else {
		rows, err = movieCollectionDB.GetMovieCollections(ctx, r.Database)
	}

	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	collections, err := movieCollectionDB.ScanConnections(rows)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	result := make([]models.Collection, 0, len(collections))
	for _, collection := range collections {
		result = append(result, collection)
	}

	return result, nil
}
