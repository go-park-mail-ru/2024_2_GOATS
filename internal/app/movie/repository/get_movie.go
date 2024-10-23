package repository

import (
	"context"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/movie"
	"github.com/rs/zerolog/log"
)

func (r *Repo) GetMovie(ctx context.Context, mvId int) (*models.MovieInfo, *errVals.ErrorObj, int) {
	rows, err := movie.FindById(ctx, mvId, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Ctx(ctx).Err(fmt.Errorf("cannot close rows while taking movie info: %w", err))
		}
	}()

	movieInfo, err := movie.ScanConnections(rows)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return movieInfo, nil, http.StatusOK
}
