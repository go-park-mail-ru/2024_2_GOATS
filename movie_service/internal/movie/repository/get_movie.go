package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/moviedb"
)

// GetMovie gets movie calling db FindByID
func (r *MovieRepo) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	rows, err := moviedb.FindByID(ctx, mvID, r.Database)
	if err != nil {
		return nil, fmt.Errorf("GetMovieRepoError: %w", err)
	}

	movieInfo, err := moviedb.ScanMovieConnection(rows)
	if err != nil {
		return nil, fmt.Errorf("GetMovieRepoError: %w", err)
	}

	genreRows, err := moviedb.GetGenres(ctx, mvID, r.Database)
	if err != nil {
		return nil, fmt.Errorf("GetMovieRepoError: %w", err)
	}

	genres, err := moviedb.ScanGenreConnection(genreRows)
	if err != nil {
		return nil, fmt.Errorf("GetMovieRepoError: %w", err)
	}

	movieInfo.Genres = genres

	return movieInfo, nil
}
