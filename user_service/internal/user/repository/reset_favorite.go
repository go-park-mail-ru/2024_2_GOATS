package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
)

func (r *UserRepo) ResetFavorite(ctx context.Context, favData *dto.RepoFavorite) error {
	err := favoritedb.Destroy(ctx, favData, r.Database)
	if err != nil {
		return fmt.Errorf("%s: %w", errors.ErrResetFavorite, err)
	}

	return nil
}
