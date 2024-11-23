package repository

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/moviedb"
)

func (r *MovieRepo) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.RepoError) {
	rows, err := moviedb.FindByID(ctx, mvID, r.Database)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	movieInfo, err := moviedb.ScanMovieConnection(rows)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	return movieInfo, nil
}
