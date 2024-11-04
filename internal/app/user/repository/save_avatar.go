package repository

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func (u *UserRepo) SaveUserAvatar(ctx context.Context, usrData *models.User) (string, *errVals.ErrorObj) {
	lclStrg := config.FromLocalStorageContext(ctx)
	fullPath := lclStrg.UserAvatarsFullUrl + usrData.AvatarName
	relativePath := lclStrg.UserAvatarsRelativeUrl + usrData.AvatarName

	outFile, fileErr := os.Create(fullPath)
	if fileErr != nil {
		return "", errVals.NewErrorObj(
			errVals.ErrFileUploadCode,
			errVals.CustomError{Err: fmt.Errorf("cannot find or create nginx static folder: %w", fileErr)},
		)
	}

	defer func() {
		if err := outFile.Close(); err != nil {
			log.Ctx(ctx).Err(fmt.Errorf("failed to close outFile: %w", err))
		}
	}()

	_, fileErr = io.Copy(outFile, usrData.Avatar)
	if fileErr != nil {
		return "", errVals.NewErrorObj(
			errVals.ErrFileUploadCode,
			errVals.CustomError{Err: fmt.Errorf("cannot save file into nginx static folder: %w", fileErr)},
		)
	}

	return relativePath, nil
}
