package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
)

func (r *UserRepo) DestroyFavorite(ctx context.Context, favData *dto.RepoFavorite) error {
	err := favoritedb.Destroy(ctx, favData, r.Database)
	if err != nil {
		return err
	}

	return nil
}
