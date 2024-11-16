package repository

import (
	"context"
	"errors"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
	"github.com/lib/pq"
)

func (u *UserRepo) UpdateProfileData(ctx context.Context, profileData *dto.DBUser) *errVals.RepoError {
	err := userdb.UpdateProfile(ctx, profileData, u.Database)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == errVals.DuplicateErrKey {
			return errVals.NewRepoError(
				errVals.DuplicateErrKey,
				errVals.NewCustomError(fmt.Sprintf("error updating profile: %v", err)),
			)
		}

		return errVals.NewRepoError(
			errVals.ErrUpdateProfileCode,
			errVals.NewCustomError(fmt.Sprintf("error updating profile: %v", err)),
		)
	}

	return nil
}
