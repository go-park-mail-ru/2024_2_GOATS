package service

import (
	"context"
	"fmt"
	"strings"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// AddFavorite sets new user favorite by calling userClient SetFavorite
func (u *UserService) AddFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError {
	err := u.userClient.SetFavorite(ctx, favData)
	if err != nil {
		if strings.Contains(err.Error(), errVals.DuplicateErrCode) {
			return errVals.NewServiceError(errVals.DuplicateErrCode, fmt.Errorf("failed to register: %w", err))
		}
		return errVals.NewServiceError("failed_to_set_favorite", err)
	}

	return nil
}
