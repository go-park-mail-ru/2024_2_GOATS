package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/userdb"
)

// UpdatePassword updates user password by calling db UpdatePassword
func (u *UserRepo) UpdatePassword(ctx context.Context, usrID uint64, pass string) error {
	hashedPasswd, err := password.HashAndSalt(ctx, pass)
	if err != nil {
		return fmt.Errorf("%s: %w", errors.ErrUpdatePasswordCode, err)
	}

	err = userdb.UpdatePassword(ctx, usrID, hashedPasswd, u.Database)
	if err != nil {
		return fmt.Errorf("%s: %w", errors.ErrUpdatePasswordCode, err)
	}

	return nil
}
