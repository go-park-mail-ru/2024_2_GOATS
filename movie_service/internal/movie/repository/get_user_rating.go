package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *MovieRepo) GetUserRating(ctx context.Context, userId int, movieId int) (float32, error) {
	var rating float32
	query := `SELECT rating FROM ratings WHERE user_id = $1 AND movie_id = $2`
	err := r.Database.QueryRowContext(ctx, query, userId, movieId).Scan(&rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("repository.GetUserRating: %w", err)
	}
	return rating, nil
}
