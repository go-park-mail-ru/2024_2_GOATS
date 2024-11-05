package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movieCollectionDB "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/movie_collectiondb"
	"github.com/rs/zerolog/log"
)

func (r *MovieRepo) GetCollection(ctx context.Context) ([]models.Collection, *errVals.ErrorObj, int) {
	rows, err := movieCollectionDB.GetMovieCollections(ctx, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Ctx(ctx).Err(fmt.Errorf("cannot close rows while taking collections: %w", err))
		}
	}()

	collections, err := movieCollectionDB.ScanConnections(rows)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	result := make([]models.Collection, 0, len(collections))
	for _, collection := range collections {
		result = append(result, collection)
	}

	return result, nil, http.StatusOK
}
