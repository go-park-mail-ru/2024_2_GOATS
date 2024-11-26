package repository

import (
	"context"
	"database/sql"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movieCollectionDB "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/movie_collectiondb"
)

func (r *MovieRepo) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.RepoError) {
	var rows *sql.Rows
	var err error

	rows, err = movieCollectionDB.GetMovieByGenre(ctx, genre, r.Database)

	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	movies, err := movieCollectionDB.ScanMovieShortInfo(rows)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	return movies, nil
}
