package repository

import (
	"context"
	"fmt"
	"log"
)

func (r *MovieRepo) DeleteUserRating(ctx context.Context, userID, movieID int) error {
	log.Println("userIDmovieID", userID, movieID)
	query := `DELETE FROM ratings WHERE user_id = $1 AND movie_id = $2`
	_, err := r.Database.ExecContext(ctx, query, userID, movieID)
	if err != nil {
		return fmt.Errorf("movie repo: failed to delete rating: %w", err)
	}
	return nil
}