package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/moviedb"
)

func (r *MovieRepo) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	rows, err := moviedb.FindByID(ctx, mvID, r.Database)
	if err != nil {
		return nil, fmt.Errorf("GetMovieRepoError: %w", err)
	}

	movieInfo, err := moviedb.ScanMovieConnection(rows)
	if err != nil {
		return nil, fmt.Errorf("GetMovieRepoError: %w", err)
	}

	return movieInfo, nil
}