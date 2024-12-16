package repository

import (
	"context"
	"fmt"
)

const (
	addOrUpdateRatingUpdateQuery = `UPDATE ratings SET rating = $1 WHERE user_id = $2 AND movie_id = $3`
	addOrUpdateRatingInsertQuery = `INSERT INTO ratings (user_id, movie_id, rating) VALUES ($1, $2, $3)`
)

// AddOrUpdateRating добавление рейтинга
func (r *MovieRepo) AddOrUpdateRating(ctx context.Context, userID int, movieID int, rating float32) error {
	res, err := r.Database.ExecContext(ctx, addOrUpdateRatingUpdateQuery, rating, userID, movieID)
	if err != nil {
		return fmt.Errorf("repository.AddOrUpdateRating (update): %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository.AddOrUpdateRating: %w", err)
	}
	if rowsAffected == 0 {
		_, err = r.Database.ExecContext(ctx, addOrUpdateRatingInsertQuery, userID, movieID, rating)
		if err != nil {
			return fmt.Errorf("repository.AddOrUpdateRating (insert): %w", err)
		}
	}

	return nil
}
