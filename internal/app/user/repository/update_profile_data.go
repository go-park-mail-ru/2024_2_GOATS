package repository

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) UpdateProfileData(ctx context.Context, profileData *dto.RepoUser) *errVals.RepoError {
	err := userdb.UpdateProfile(ctx, profileData, u.Database)

	if err != nil {
		if errVals.IsDuplicateError(err) {
			errMsg := fmt.Sprintf("error updating profile: %v", err)
			return errVals.NewRepoError(errVals.DuplicateErrCode, errVals.NewCustomError(errMsg))
		}

		return errVals.NewRepoError(
			errVals.ErrUpdateProfileCode,
			errVals.NewCustomError(fmt.Sprintf("error updating profile: %v", err)),
		)
	}

	return nil
}
