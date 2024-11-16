package repository

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/favoritedb"
)

func (r *UserRepo) DestroyFavorite(ctx context.Context, favData *dto.DBFavorite) *errVals.RepoError {
	err := favoritedb.Destroy(ctx, favData, r.Database)
	if err != nil {
		return errVals.NewRepoError(errVals.ErrDestroyFavorite, errVals.NewCustomError(err.Error()))
	}

	return nil
}
