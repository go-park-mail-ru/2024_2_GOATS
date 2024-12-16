package repository

import (
	"context"
	"fmt"
)

const DeleteUserRatingQuery = `DELETE FROM ratings WHERE user_id = $1 AND movie_id = $2`

func (r *MovieRepo) DeleteUserRating(ctx context.Context, userID, movieID int) error {
	_, err := r.Database.ExecContext(ctx, DeleteUserRatingQuery, userID, movieID)
	if err != nil {
		return fmt.Errorf("movie repo: failed to delete rating: %w", err)
	}
	return nil
}
