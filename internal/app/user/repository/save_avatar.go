package repository

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserRepo) SaveAvatar(ctx context.Context, usrData *models.User) (string, *errVals.ErrorObj) {
	locS := config.FromContext(ctx).Databases.LocalStorage
	fullPath := locS.UserAvatarsFullUrl + usrData.AvatarName
	relativePath := locS.UserAvatarsRelativeUrl + usrData.AvatarName

	outFile, osErr := os.Create(fullPath)
	if osErr != nil {
		return "", &errVals.ErrorObj{
			Code:  "file_upload_err",
			Error: errVals.CustomError{Err: fmt.Errorf("cannot find or create nginx static folder: %w", osErr)},
		}
	}

	defer outFile.Close()

	_, osErr = io.Copy(outFile, usrData.Avatar)
	if osErr != nil {
		return "", &errVals.ErrorObj{
			Code:  "file_upload_err",
			Error: errVals.CustomError{Err: fmt.Errorf("cannot save file into nginx static folder: %w", osErr)},
		}
	}

	return relativePath, nil
}
