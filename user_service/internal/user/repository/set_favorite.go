package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
)

// SetFavorite creates new user favorites relation by calling db CreateFavorite
func (r *UserRepo) SetFavorite(ctx context.Context, favData *dto.RepoFavorite) error {
	err := favoritedb.Create(ctx, favData, r.Database)
	if err != nil {
		if errors.IsDuplicateError(err) {
			errMsg := fmt.Errorf("%s: %w", errors.ErrCreateFavorite, err)
			return fmt.Errorf("%s: %w", errors.DuplicateErrCode, errMsg)
		}

		return fmt.Errorf("%s: %w", errors.ErrCreateFavorite, err)
	}

	return nil
}
