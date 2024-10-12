package repository

import (
	"context"
	"fmt"
	"net/http"
)

func (r *Repo) GetFromCookie(ctx context.Context, cookie string) (string, error, int) {
	var userID string
	err := r.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("cannot get cookie from redis: %w", err), http.StatusInternalServerError
	}

	return userID, nil, http.StatusOK
}
