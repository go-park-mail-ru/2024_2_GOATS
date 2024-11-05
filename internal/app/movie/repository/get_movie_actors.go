package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/moviedb"
	"github.com/rs/zerolog/log"
)

func (r *MovieRepo) GetMovieActors(
	ctx context.Context,
	mvID int,
) ([]*models.ActorInfo, *errVals.ErrorObj, int) {
	logger := log.Ctx(ctx)
	rows, err := moviedb.GetMovieActors(ctx, mvID, r.Database)

	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Error().Err(err).Msg("cannot close rows while taking movie actors")
		}
	}()

	actorsInfos, err := moviedb.ScanActorsConnections(rows)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return actorsInfos, nil, http.StatusOK
}
