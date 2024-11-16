package repository

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) UpdatePassword(ctx context.Context, usrID int, pass string) *errVals.RepoError {
	hashedPasswd, err := password.HashAndSalt(ctx, pass)
	if err != nil {
		return errVals.NewRepoError(
			errVals.ErrServerCode,
			errVals.NewCustomError(fmt.Sprintf("error hashing password: %v", err)),
		)
	}

	err = userdb.UpdatePassword(ctx, usrID, hashedPasswd, u.Database)
	if err != nil {
		return errVals.NewRepoError(
			errVals.ErrUpdatePasswordCode,
			errVals.NewCustomError(fmt.Sprintf("error updating password: %v", err)),
		)
	}

	return nil
}
