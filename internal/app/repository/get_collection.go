package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movieCollection "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/movie_collection"
	"github.com/labstack/gommon/log"
)

func (r *Repo) GetCollection(ctx context.Context) ([]models.Collection, *errVals.ErrorObj, int) {
	rows, err := movieCollection.Obtain(ctx, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Errorf("cannot close rows while taking collections: %w", err)
		}
	}()

	collections, err := movieCollection.ScanConnections(rows)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	result := make([]models.Collection, 0, len(collections))
	for _, collection := range collections {
		result = append(result, collection)
	}

	return result, nil, http.StatusOK
}
