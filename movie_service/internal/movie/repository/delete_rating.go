package repository

import (
	"context"
	"fmt"
)

const deleteUserRatingQuery = `DELETE FROM ratings WHERE user_id = $1 AND movie_id = $2`

// DeleteUserRating удаление рейтинга
func (r *MovieRepo) DeleteUserRating(ctx context.Context, userID, movieID int) error {
	_, err := r.Database.ExecContext(ctx, deleteUserRatingQuery, userID, movieID)
	if err != nil {
		return fmt.Errorf("movie repo: failed to delete rating: %w", err)
	}
	return nil
}
