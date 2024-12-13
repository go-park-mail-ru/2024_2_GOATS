package repository

import (
	"context"
	"fmt"
)

func (r *MovieRepo) UpdateMovieRating(ctx context.Context, movieId int) error {
	query := `
		UPDATE movies 
		SET rating = (
			SELECT COALESCE(AVG(rating), 0) 
			FROM ratings 
			WHERE movie_id = $1
		)
		WHERE id = $1
	`
	_, err := r.Database.ExecContext(ctx, query, movieId)
	if err != nil {
		return fmt.Errorf("repository.UpdateMovieRating: %w", err)
	}
	return nil
}
