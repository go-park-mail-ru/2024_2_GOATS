package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/actordb"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/moviedb"
)

func (r *MovieRepo) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ErrorObj, int) {
	actor, err := actordb.FindByID(ctx, actorID, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	rows, err := moviedb.FindByActorID(ctx, actorID, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	actMvs, err := moviedb.ScanActorMoviesConnections(rows)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	actor.Movies = actMvs

	return actor, nil, http.StatusOK
}
