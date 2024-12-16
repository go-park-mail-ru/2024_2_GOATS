package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/actordb"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/moviedb"
)

// GetActor get actor using db FindByID
func (r *MovieRepo) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	actor, err := actordb.FindByID(ctx, actorID, r.Database)
	if err != nil {
		return nil, fmt.Errorf("getActorRepoError: %w", err)
	}

	rows, err := moviedb.FindByActorID(ctx, actorID, r.Database)
	if err != nil {
		return nil, fmt.Errorf("getActorRepoError: %w", err)
	}

	actMvs, err := moviedb.ScanActorMoviesConnections(rows)
	if err != nil {
		return nil, fmt.Errorf("getActorRepoError: %w", err)
	}

	srvAct := converter.ToActorInfoFromRepo(actor)
	for _, mv := range actMvs {
		srvAct.Movies = append(srvAct.Movies, converter.ToMovieShortInfoFromRepo(mv))
	}

	return srvAct, nil
}
