package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const getUserRatingQuery = `SELECT rating FROM ratings WHERE user_id = $1 AND movie_id = $2`

// GetUserRating gets user_rating for movie
func (r *MovieRepo) GetUserRating(ctx context.Context, userID int, movieID int) (float32, error) {
	var rating float32
	err := r.Database.QueryRowContext(ctx, getUserRatingQuery, userID, movieID).Scan(&rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("repository.GetUserRating: %w", err)
	}
	return rating, nil
}
