package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/moviedb"
	"github.com/rs/zerolog/log"
)

// GetMovieActors gets actors for a movie
func (r *MovieRepo) GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error) {
	logger := log.Ctx(ctx)
	rows, err := moviedb.GetMovieActors(ctx, mvID, r.Database)

	if err != nil {
		return nil, fmt.Errorf("getMovieActorsRepoError: %w", err)
	}

	defer func() {
		if clErr := rows.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("cannot close rows while taking movie_service actors")
		}
	}()

	actorsInfos, err := moviedb.ScanActorsConnections(rows)
	if err != nil {
		return nil, fmt.Errorf("getMovieActorsRepoError: %w", err)
	}

	var srvActors = make([]*models.ActorInfo, 0, len(actorsInfos))
	for _, ac := range actorsInfos {
		srvActors = append(srvActors, converter.ToActorInfoFromRepo(ac))
	}

	return srvActors, nil
}
