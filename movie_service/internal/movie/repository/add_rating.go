package repository

import (
	"context"
	"fmt"
)

func (r *MovieRepo) AddOrUpdateRating(ctx context.Context, userId int, movieId int, rating float32) error {
	query := `UPDATE ratings SET rating = $1 WHERE user_id = $2 AND movie_id = $3`
	res, err := r.Database.ExecContext(ctx, query, rating, userId, movieId)
	if err != nil {
		return fmt.Errorf("repository.AddOrUpdateRating (update): %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository.AddOrUpdateRating (check rows affected): %w", err)
	}
	if rowsAffected == 0 {
		query = `INSERT INTO ratings (user_id, movie_id, rating) VALUES ($1, $2, $3)`
		_, err = r.Database.ExecContext(ctx, query, userId, movieId, rating)
		if err != nil {
			return fmt.Errorf("repository.AddOrUpdateRating (insert): %w", err)
		}
	}

	return nil
}
