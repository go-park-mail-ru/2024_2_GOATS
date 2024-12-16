package repository

import (
	"context"
	"fmt"
)

const UpdateMovieRatingQuery = `
		UPDATE movies 
		SET rating = (
			SELECT COALESCE(AVG(rating), 0) 
			FROM ratings 
			WHERE movie_id = $1
		)
		WHERE id = $1
	`

func (r *MovieRepo) UpdateMovieRating(ctx context.Context, movieId int) error {
	_, err := r.Database.ExecContext(ctx, UpdateMovieRatingQuery, movieId)
	if err != nil {
		return fmt.Errorf("repository.UpdateMovieRating: %w", err)
	}
	return nil
}
