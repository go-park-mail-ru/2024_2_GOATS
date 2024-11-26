package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	movieCollectionDB "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/movie_collectiondb"
)

func (r *MovieRepo) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	var rows *sql.Rows
	var err error

	rows, err = movieCollectionDB.GetMovieByGenre(ctx, genre, r.Database)

	if err != nil {
		return nil, fmt.Errorf("GetMovieByGenreRepoError: %w", err)
	}

	movies, err := movieCollectionDB.ScanMovieShortInfo(rows)
	if err != nil {
		return nil, fmt.Errorf("GetMovieByGenreRepoError: %w", err)
	}

	return movies, nil
}
