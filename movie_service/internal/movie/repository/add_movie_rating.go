package repository

import (
	"context"
	"fmt"
)

const updateMovieRatingQuery = `
		UPDATE movies 
		SET rating = (
			SELECT COALESCE(AVG(rating), 0) 
			FROM ratings 
			WHERE movie_id = $1
		)
		WHERE id = $1
	`

// UpdateMovieRating обновление рейтинга
func (r *MovieRepo) UpdateMovieRating(ctx context.Context, movieID int) error {
	_, err := r.Database.ExecContext(ctx, updateMovieRatingQuery, movieID)
	if err != nil {
		return fmt.Errorf("repository.UpdateMovieRating: %w", err)
	}
	return nil
}
