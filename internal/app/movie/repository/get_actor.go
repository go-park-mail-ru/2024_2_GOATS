package repository

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/actordb"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/moviedb"
)

func (r *MovieRepo) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.RepoError) {
	actor, err := actordb.FindByID(ctx, actorID, r.Database)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	rows, err := moviedb.FindByActorID(ctx, actorID, r.Database)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	actMvs, err := moviedb.ScanActorMoviesConnections(rows)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	srvAct := converter.ToActorInfoFromRepo(actor)
	for _, mv := range actMvs {
		srvAct.Movies = append(srvAct.Movies, converter.ToMovieShortInfoFromRepo(mv))
	}

	return srvAct, nil
}
