package repository

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
)

func (r *UserRepo) DestroyFavorite(ctx context.Context, favData *dto.RepoFavorite) *errVals.RepoError {
	err := favoritedb.Destroy(ctx, favData, r.Database)
	if err != nil {
		return errVals.NewRepoError(errVals.ErrResetFavorite, errVals.NewCustomError(err.Error()))
	}

	return nil
}
