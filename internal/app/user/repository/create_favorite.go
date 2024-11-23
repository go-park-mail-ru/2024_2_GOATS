package repository

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/favoritedb"
)

func (r *UserRepo) CreateFavorite(ctx context.Context, favData *dto.RepoFavorite) *errVals.RepoError {
	err := favoritedb.Create(ctx, favData, r.Database)
	if err != nil {
		if errVals.IsDuplicateError(err) {
			errMsg := fmt.Sprintf("error creating favorite: %v", err)
			return errVals.NewRepoError(errVals.DuplicateErrCode, errVals.NewCustomError(errMsg))
		}

		return errVals.NewRepoError(errVals.ErrCreateFavorite, errVals.NewCustomError(err.Error()))
	}

	return nil
}
