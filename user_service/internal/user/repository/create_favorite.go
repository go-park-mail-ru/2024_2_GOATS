package repository

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
)

func (r *UserRepo) CreateFavorite(ctx context.Context, favData *dto.RepoFavorite) error {
	err := favoritedb.Create(ctx, favData, r.Database)
	if err != nil {
		if errVals.IsDuplicateError(err) {
			errMsg := fmt.Errorf("error creating favorite: %s : %w", errVals.DuplicateErrCode, err)
			return errMsg
		}

		return err
	}

	return nil
}
