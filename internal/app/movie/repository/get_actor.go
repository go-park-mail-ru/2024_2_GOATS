package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/actor"
)

func (r *Repo) GetActor(ctx context.Context, actorId int) (*models.ActorInfo, *errVals.ErrorObj, int) {
	actor, err := actor.FindById(ctx, actorId, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return actor, nil, http.StatusOK
}
