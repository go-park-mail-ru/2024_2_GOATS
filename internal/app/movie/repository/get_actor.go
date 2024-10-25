package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/staff/actor"
)

func (r *Repo) GetActor(ctx context.Context, actorId int) (*models.StaffInfo, *errVals.ErrorObj, int) {
	actor, err := actor.FindById(ctx, actorId, "actor", r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return actor, nil, http.StatusOK
}
