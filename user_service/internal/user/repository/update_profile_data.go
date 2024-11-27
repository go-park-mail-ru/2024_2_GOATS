package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/userdb"
)

func (u *UserRepo) UpdateProfileData(ctx context.Context, profileData *dto.RepoUser) error {
	err := userdb.UpdateProfile(ctx, profileData, u.Database)

	if err != nil {
		if errors.IsDuplicateError(err) {
			errMsg := fmt.Errorf("%s: %w", errors.ErrUpdateProfileCode, err)
			return fmt.Errorf("%s: %w", errors.DuplicateErrCode, errMsg)
		}

		return fmt.Errorf("%s: %w", errors.ErrUpdateProfileCode, err)
	}

	return nil
}
