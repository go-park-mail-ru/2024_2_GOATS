package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
)

// CheckFavorite checks if user has movie in favorites by calling db Check
func (r *UserRepo) CheckFavorite(ctx context.Context, favData *dto.RepoFavorite) (bool, error) {
	present, err := favoritedb.Check(ctx, favData, r.Database)
	if err != nil {
		return false, fmt.Errorf("%s: %w", errors.ErrCreateFavorite, err)
	}

	return present, nil
}
