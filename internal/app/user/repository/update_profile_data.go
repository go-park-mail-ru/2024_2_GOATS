package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
	"github.com/lib/pq"
)

func (u *UserRepo) UpdateProfileData(ctx context.Context, profileData *models.User) (*errVals.ErrorObj, int) {
	err := userdb.UpdateProfile(ctx, profileData, u.Database)

	if err != nil {
		status := http.StatusInternalServerError
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == errVals.DuplicateErrKey {
			status = http.StatusConflict
		}

		return errVals.NewErrorObj(
			errVals.ErrUpdateProfileCode,
			errVals.CustomError{Err: fmt.Errorf("error updating profile: %w", err)},
		), status
	}

	return nil, http.StatusOK
}
