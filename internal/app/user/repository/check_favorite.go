package repository

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/favoritedb"
)

func (r *UserRepo) CheckFavorite(ctx context.Context, favData *dto.RepoFavorite) (bool, *errVals.RepoError) {
	present, err := favoritedb.Check(ctx, favData, r.Database)
	if err != nil {
		return false, errVals.NewRepoError(errVals.ErrCreateFavorite, errVals.NewCustomError(err.Error()))
	}

	return present, nil
}
