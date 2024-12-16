package repository

import (
	"context"
	"fmt"
)

const (
	AddOrUpdateRatingUpdateQuery = `UPDATE ratings SET rating = $1 WHERE user_id = $2 AND movie_id = $3`
	AddOrUpdateRatingInsertQuery = `INSERT INTO ratings (user_id, movie_id, rating) VALUES ($1, $2, $3)`
)

func (r *MovieRepo) AddOrUpdateRating(ctx context.Context, userId int, movieId int, rating float32) error {
	res, err := r.Database.ExecContext(ctx, AddOrUpdateRatingUpdateQuery, rating, userId, movieId)
	if err != nil {
		return fmt.Errorf("repository.AddOrUpdateRating (update): %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository.AddOrUpdateRating: %w", err)
	}
	if rowsAffected == 0 {
		_, err = r.Database.ExecContext(ctx, AddOrUpdateRatingInsertQuery, userId, movieId, rating)
		if err != nil {
			return fmt.Errorf("repository.AddOrUpdateRating (insert): %w", err)
		}
	}

	return nil
}
