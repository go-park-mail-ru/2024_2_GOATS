package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const getUserRatingQuery = `SELECT rating FROM ratings WHERE user_id = $1 AND movie_id = $2`

func (r *MovieRepo) GetUserRating(ctx context.Context, userId int, movieId int) (float32, error) {
	var rating float32
	err := r.Database.QueryRowContext(ctx, getUserRatingQuery, userId, movieId).Scan(&rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("repository.GetUserRating: %w", err)
	}
	return rating, nil
}
